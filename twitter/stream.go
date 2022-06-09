package twitter

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// Stream connects to a streaming endpoint on the Twitter API.
// It receives messages from the streaming endpoint and sends them on the
// Queue channel from a goroutine.
type Stream struct {
	Config       Config
	APIInfo      APIInfo
	EndPointInfo *EndPointInfo
	AttemptTime  time.Time
	Time         time.Time
	HTTPRequest  *http.Request
	HTTPResponse *http.Response
	Retryable    *bool
	PayLoad      interface{}
	Error        error
	Data         interface{}
	Handlers     StreamHandlers
	Retryer
	MessageQueue chan interface{}
	rawData      chan []byte
	done         chan struct{}
	errorChan    chan error
	waitGroup    *sync.WaitGroup
	body         io.ReadCloser
}

// NewStream returns a new stream object that is connect to a streaming endpoint.
// If there is no error in the stream connection you can start reading messages
// from the Queue channel
func (c *Client) NewStream(endpoint *EndPointInfo, input, output interface{}) *Stream {
	return createStream(*c.Config, c.APIInfo, c.Retryer, endpoint, input, output)
}

func createStream(cfg Config, apiInfo APIInfo,
	retryer Retryer, endpointInfo *EndPointInfo, payLoad interface{}, data interface{}) *Stream {
	var err error

	if retryer == nil {
		retryer = noRetryer{}
	}

	if err = endpointInfo.Validate(); err != nil {
		return &Stream{Error: err}
	}

	httpReq, err := http.NewRequest(endpointInfo.HTTPMethod, "", nil)
	if err != nil {
		return &Stream{Error: err}
	}

	httpReq.Header.Add("Content-type", "application/json")

	httpReq.URL, err = url.Parse(apiInfo.Endpoint + "/" + apiInfo.APIVersion + "/" + endpointInfo.HTTPPath)
	if err != nil {
		httpReq.URL = &url.URL{}
		return &Stream{Error: err}
	}

	if endpointInfo.QueryParams != nil {
		q := httpReq.URL.Query()
		for k, v := range endpointInfo.QueryParams {
			q.Add(k, v)
		}
		httpReq.URL.RawQuery = q.Encode()
	}

	if payLoad != nil {
		b, err := json.Marshal(payLoad)
		if err != nil {
			return &Stream{Error: err}
		}
		httpReq.Body = ioutil.NopCloser(strings.NewReader(string(b)))
	}

	handlers := StreamHandlers{
		Send: StreamSendHandler,
		Sign: StreamSigner,
	}

	s := &Stream{
		Config:       cfg,
		APIInfo:      apiInfo,
		EndPointInfo: endpointInfo,
		Handlers:     handlers.Copy(),

		Retryer:      retryer,
		Time:         time.Now(),
		HTTPRequest:  httpReq,
		waitGroup:    &sync.WaitGroup{},
		PayLoad:      payLoad,
		Error:        err,
		Data:         data,
		MessageQueue: make(chan interface{}),
		rawData:      make(chan []byte),
		done:         make(chan struct{}),
	}
	s.waitGroup.Add(2)
	go s.consume()
	go s.processMessage()

	return s

}

func (s *Stream) consume() {

	defer close(s.rawData)

	defer s.waitGroup.Done()
	for !s.stopped() {
		s.Error = nil
		s.AttemptTime = time.Now()

		if err := s.sign(); err != nil {
			s.Config.Logger.Error().Err(err).Msg("Failed to sign stream request")
			return
		}

		if err := s.sendRequest(); err == nil {
			s.receive(s.body)
		}

		s.Handlers.Retry.Run(s)

		if s.Error != nil || !BoolValue(s.Retryable) {
			s.Config.Logger.Error().Err(s.Error).Msg("stream request can not be retried")
			return
		}

	}

}

// Sign will sign the request, returning error if errors are encountered.
func (s *Stream) sign() error {
	s.Handlers.Sign.Run(s)
	return s.Error
}

func (s *Stream) sendRequest() (sendErr error) {
	s.Retryable = nil
	s.Handlers.Send.Run(s)
	if s.Error != nil {
		return s.Error
	}
	s.body = s.HTTPResponse.Body
	return nil

}

func (s *Stream) stopped() bool {
	select {
	case <-s.done:
		return true
	default:
		return false
	}
}

// Stop signals retry and receiver to stop, closes the Messages channel, and
// blocks until done.
func (s *Stream) Stop() {
	close(s.done)
	// Scanner does not have a Stop() or take a done channel, so for low volume
	// streams Scan() blocks until the next keep-alive. Close the resp.Body to
	// escape and stop the stream in a timely fashion.
	if s.body != nil {
		s.body.Close()
	}
	// block until the retry goroutine stops
	s.waitGroup.Wait()
}

func (s *Stream) receive(body io.Reader) {
	reader := newStreamResponseBodyReader(body)
	for !s.stopped() {
		data, err := reader.readNext()
		if err != nil {
			message := "failed to read next tweet from streaming response body"
			s.Config.Logger.Error().Err(err).Msg(message)
			return
		}
		if len(data) == 0 {
			// empty keep-alive
			continue
		}

		select {
		// send messages, data, or errors
		case s.rawData <- data:
			continue

		// allow client to Stop(), even if not receiving
		case <-s.done:
			return
		}
	}
}

func (s *Stream) processMessage() {
	defer close(s.MessageQueue)
	defer s.waitGroup.Done()
	for !s.stopped() {
		messageBytes, ok := <-s.rawData
		if !ok {
			return
		}
		message, err := s.getMessage(messageBytes)
		if err != nil {
			message := "Failed to get message from raw data"
			s.Config.Logger.Error().Err(err).Msg(message)
			s.Retryable = Bool(false)
			return
		}

		select {
		// send messages, data, or errors
		case s.MessageQueue <- message:
			continue

		// allow client to Stop(), even if not receiving
		case <-s.done:
			return
		}
	}

}

func (s *Stream) getMessage(messageBytes []byte) (interface{}, error) {
	reader := bytes.NewReader(messageBytes)
	err := UnmarshalJSON(s.Data, reader)
	if err != nil {
		message := "Failed to unmarshal bytes in json"
		s.Config.Logger.Error().Err(err).Msg(message)
		return nil, err
	}
	return s.Data, nil
}
