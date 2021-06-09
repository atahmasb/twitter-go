package twitter

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

// A Request is the service request to be made.
type Request struct {
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
	Handlers     Handlers
	Retryer
}

// An EndPointInfo is the endpoint info to create the request.
type EndPointInfo struct {
	Name        string
	HTTPMethod  string
	HTTPPath    string
	QueryParams map[string]string
}

// Validate checks if an endpoint struct has required fields before being used in sending a request
func (e *EndPointInfo) Validate() error {
	if e.Name == "" || e.HTTPMethod == "" || e.HTTPPath == "" {
		return errors.New("endpoint info is missing required fields")
	}
	return nil
}

func (e *EndPointInfo) String() string {
	return fmt.Sprintf("name: %s, method: %s, path: %s", e.Name, e.HTTPMethod, e.HTTPPath)
}

// NewRequest returns a new Request pointer for the service API
// operation and parameters.
func (c *Client) NewRequest(endpoint *EndPointInfo, input, output interface{}) *Request {
	return CreateRequest(*c.Config, c.APIInfo, c.Retryer, endpoint, input, output)
}

// CreateRequest returns a new Request pointer for the service API
// operation and parameters.
//
// A Retryer should be provided to direct how the request is retried. If
// Retryer is nil, a default no retry value will be used. You can use
// NoOpRetryer in the Client package to disable retry behavior directly.
//
// Params is any value of input parameters to be the request payload.
// Data is pointer value to an object which the request's response
// payload will be deserialized to.
func CreateRequest(cfg Config, apiInfo APIInfo,
	retryer Retryer, endpointInfo *EndPointInfo, payLoad interface{}, data interface{}) *Request {
	var err error

	if retryer == nil {
		retryer = noRetryer{}
	}

	if err = endpointInfo.Validate(); err != nil {
		return &Request{Error: err}
	}

	httpReq, err := http.NewRequest(endpointInfo.HTTPMethod, "", nil)
	if err != nil {
		return &Request{Error: err}
	}

	httpReq.Header.Add("Content-type", "application/json")

	httpReq.URL, err = url.Parse(apiInfo.Endpoint + "/" + apiInfo.APIVersion + "/" + endpointInfo.HTTPPath)
	if err != nil {
		httpReq.URL = &url.URL{}
		return &Request{Error: err}
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
			return &Request{Error: err}
		}
		httpReq.Body = ioutil.NopCloser(strings.NewReader(string(b)))
	}

	handlers := Handlers{
		Send:           SendHandler,
		Unmarshal:      UnMarshaler,
		ErrorUnmarshal: ErrorUnmarshaler,
		Sign:           Signer,
	}

	r := &Request{
		Config:       cfg,
		APIInfo:      apiInfo,
		EndPointInfo: endpointInfo,
		Handlers:     handlers.Copy(),

		Retryer:     retryer,
		Time:        time.Now(),
		HTTPRequest: httpReq,
		// Body:        nil,
		PayLoad: payLoad,
		Error:   err,
		Data:    data,
	}
	return r

}

// Send will send the request, returning error if errors are encountered.
func (r *Request) Send() error {
	if err := r.Error; err != nil {
		return err
	}

	for {
		r.Error = nil
		r.AttemptTime = time.Now()

		if err := r.Sign(); err != nil {
			return err
		}

		if err := r.sendRequest(); err == nil {
			return nil
		}
		r.Handlers.Retry.Run(r)
		r.Handlers.AfterRetry.Run(r)

		if r.Error != nil || !BoolValue(r.Retryable) {
			return r.Error
		}

	}
}

// Sign will sign the request, returning error if errors are encountered.
func (r *Request) Sign() error {
	if r.Error != nil {
		debugLogReqError(r, "Sign Request", r.Error)
	}
	r.Handlers.Sign.Run(r)
	return r.Error
}

func (r *Request) sendRequest() (sendErr error) {
	r.Retryable = nil
	r.Handlers.Send.Run(r)
	if r.Error != nil {
		debugLogReqError(r, "Send Request", r.Error)
		return r.Error
	}

	r.Handlers.ErrorUnmarshal.Run(r)
	if r.Error != nil {
		debugLogReqError(r, "Unmarshal Response Errors", r.Error)
		return r.Error
	}

	r.Handlers.Unmarshal.Run(r)
	if r.Error != nil {
		debugLogReqError(r, "Unmarshal Response", r.Error)
		return r.Error
	}

	return nil
}

func debugLogReqError(r *Request, stage string, err error) {
	r.Config.Logger.Debug().Str("stage", stage).Msg(err.Error())
}

// DataFilled returns true if the request's data for response deserialization
// target has been set and is a valid. False is returned if data is not
// set, or is invalid.
func (r *Request) DataFilled() bool {
	return r.Data != nil && reflect.ValueOf(r.Data).Elem().IsValid()
}
