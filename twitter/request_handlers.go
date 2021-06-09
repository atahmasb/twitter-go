package twitter

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// A Handlers provides a collection of request handlers for various
// stages of handling requests.
type Handlers struct {
	Sign           HandlerFunction
	Send           HandlerFunction
	Unmarshal      HandlerFunction
	ErrorUnmarshal HandlerFunction
	Retry          HandlerFunction
	AfterRetry     HandlerFunction
}

// Copy returns a copy of this handler's lists.
func (h *Handlers) Copy() Handlers {
	return Handlers{
		Sign:           h.Sign.copy(),
		Send:           h.Send.copy(),
		Unmarshal:      h.Unmarshal.copy(),
		ErrorUnmarshal: h.ErrorUnmarshal.copy(),
		Retry:          h.Retry.copy(),
		AfterRetry:     h.AfterRetry.copy(),
	}
}

// A HandlerFunction is a struct that contains a name and function callback.
type HandlerFunction struct {
	Name string
	Fn   func(*Request)
}

func (h *HandlerFunction) copy() HandlerFunction {
	n := HandlerFunction{
		Name: h.Name,
		Fn:   h.Fn,
	}
	return n
}

// Run executes callback function.
func (h HandlerFunction) Run(r *Request) {
	if h.Fn != nil {
		h.Fn(r)
	}

}

// SendHandler is a request handler to send service request using HTTP client.
var SendHandler = HandlerFunction{
	Name: "SendHandler",
	Fn: func(r *Request) {
		sender := send

		var err error
		r.HTTPResponse, err = sender(r)
		if err != nil {
			handleSendError(r, err)
		}
	},
}

// UnMarshaler is a request handler to unmarshal HTTP response to Golang struct.
var UnMarshaler = HandlerFunction{
	Name: "UnMarshaler",
	Fn: func(r *Request) {
		defer r.HTTPResponse.Body.Close()
		if r.DataFilled() {

			err := UnmarshalJSON(r.Data, r.HTTPResponse.Body)
			if err != nil {

				r.Error = NewRequestFailure(
					err,
					r.HTTPResponse.StatusCode,
					"Failed to decode JSON response",
				)
			}
		}
	},
}

// Signer is a request handler to add the credentials to the request header.
var Signer = HandlerFunction{
	Name: "Signer",
	Fn: func(r *Request) {

		token, err := r.Config.Credentials.Retrieve()
		if err != nil {
			r.Error = err
		}
		r.HTTPRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.BearerToken))
	},
}

// ErrorUnmarshaler unmarshals the API request's errors and adds
// an error to Request if errors are in HTTP Response body
var ErrorUnmarshaler = HandlerFunction{
	Name: "ErrorUnmarshaler",
	Fn: func(r *Request) {
		if r.HTTPResponse.StatusCode >= 400 {
			var diag Diagnostic
			err := UnmarshalJSON(&diag, r.HTTPResponse.Body)
			if err != nil {
				r.Error = NewRequestFailure(
					err,
					r.HTTPResponse.StatusCode,
					"Failed to decode JSON response to detect errors",
				)
			}

			if diag.IsFiled() {
				r.Error = NewRequestFailure(
					nil,
					r.HTTPResponse.StatusCode,
					fmt.Sprintf("An error ocurred in sending a request to: %s with error details: %v", r.EndPointInfo.String(), diag.String()),
				)
			}

		}
	},
}

func send(r *Request) (*http.Response, error) {
	return r.Config.HTTPClient.Do(r.HTTPRequest)
}

func handleSendError(r *Request, err error) {

	if r.HTTPResponse != nil {
		r.HTTPResponse.Body.Close()
	}

	if r.HTTPResponse == nil {
		// Add a dummy request response object to ensure the HTTPResponse
		// value is consistent.
		r.HTTPResponse = &http.Response{
			StatusCode: int(0),
			Status:     http.StatusText(int(0)),
			Body:       ioutil.NopCloser(bytes.NewReader([]byte{})),
		}
	}
	// Catch all request errors, and let the default retryer determine
	// if the error is retryable.
	r.Error = NewRequestFailure(err, r.HTTPResponse.StatusCode, "send request failed")

}
