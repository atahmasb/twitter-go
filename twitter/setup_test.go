package twitter

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/suite"
)

type twitterClientSuite struct {
	suite.Suite
	client *Client
	server *httptest.Server
	mux    *http.ServeMux
}

func Test_twitterClientSuite(t *testing.T) {
	suite.Run(t, new(twitterClientSuite))
}

func (suite *twitterClientSuite) SetupTest() {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	transport := &RewriteTransport{&http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}}
	httpClient := &http.Client{Transport: transport}

	config := NewConfig().
		WithCredentials(NewCredentials(Value{
			BearerToken: "TEST",
		})).
		WithHTTPClient(httpClient).
		WithLogLevel(int8(2)).
		WithLogger(NewDefaultLogger()).
		WithRetryer(noRetryer{})

	client := &Client{
		Config:  config,
		APIInfo: newAPIInfo(),
	}

	suite.client = client
	suite.server = server
	suite.mux = mux

}

func (suite *twitterClientSuite) Test_Close() {
	suite.server.Close()

}

type RewriteTransport struct {
	Transport http.RoundTripper
}

func (t *RewriteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Scheme = "http"
	if t.Transport == nil {
		return http.DefaultTransport.RoundTrip(req)
	}
	return t.Transport.RoundTrip(req)
}

// assertQuery tests that the Request has the expected url query key/val pairs
func (suite *twitterClientSuite) assertQuery(expected map[string]string, req *http.Request) {
	queryValues := req.URL.Query()
	expectedValues := url.Values{}
	for key, value := range expected {
		expectedValues.Add(key, value)
	}
	suite.Assert().Equal(expectedValues, queryValues)
}

// assertMethod tests that the Request has the expected http method
func (suite *twitterClientSuite) assertMethod(expectedMethod string, req *http.Request) {
	suite.Assert().Equal(expectedMethod, req.Method)
}
