package twitter

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// A StreamHandlers provides a collection of stream request handlers for various
// stages of handling requests.
type StreamHandlers struct {
	Sign           StreamHandlerFunction
	Send           StreamHandlerFunction
	Retry          StreamHandlerFunction
	AfterRetry     StreamHandlerFunction
}

// Copy returns a copy of this handler's lists.
func (h *StreamHandlers) Copy() StreamHandlers {
	return StreamHandlers{
		Sign:           h.Sign.copy(),
		Send:           h.Send.copy(),
		Retry:          h.Retry.copy(),
		AfterRetry:     h.AfterRetry.copy(),
	}
}

// A StreamHandlerFunction is a struct that contains a name and function callback.
type StreamHandlerFunction struct {
	Name string
	Fn   func(*Stream)
}

func (h *StreamHandlerFunction) copy() StreamHandlerFunction {
	n := StreamHandlerFunction{
		Name: h.Name,
		Fn:   h.Fn,
	}
	return n
}


// Run executes callback function.
func (h StreamHandlerFunction) Run(s *Stream) {
	if h.Fn != nil {
		h.Fn(s)
	}

}

// StreamSendHandler is a stream request handler to send service request using HTTP client.
var StreamSendHandler = StreamHandlerFunction{
	Name: "SendHandler",
	Fn: func(s *Stream) {
		sender := sendStreamRequest

		var err error
		s.HTTPResponse, err = sender(s)
		if err != nil {
			handleStreamSendError(s, err)
		}
	},
}



// StreamSigner is a stream handler to add the credentials to the stream request header.
var StreamSigner = StreamHandlerFunction{
	Name: "Signer",
	Fn: func(s *Stream) {

		token, err := s.Config.Credentials.Retrieve()
		if err != nil {
			s.Error = err
		}
		s.HTTPRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.BearerToken))
	},
}

func sendStreamRequest(s *Stream) (*http.Response, error) {
	return s.Config.HTTPClient.Do(s.HTTPRequest)
}

func handleStreamSendError(s *Stream, err error) {

	if s.HTTPResponse != nil {
		s.HTTPResponse.Body.Close()
	}

	if s.HTTPResponse == nil {
		// Add a dummy request response object to ensure the HTTPResponse
		// value is consistent.
		s.HTTPResponse = &http.Response{
			StatusCode: int(0),
			Status:     http.StatusText(int(0)),
			Body:       ioutil.NopCloser(bytes.NewReader([]byte{})),
		}
	}
	// Catch all request errors, and let the default retryer determine
	// if the error is retryable.
	s.Error = NewRequestFailure(err, s.HTTPResponse.StatusCode, "send request failed")

}


