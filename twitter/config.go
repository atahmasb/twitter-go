package twitter

import (
	"net/http"
	"os"

	"github.com/rs/zerolog"
)

// RequestRetryer is an alias for a type that implements the Retryer interface.
type RequestRetryer interface{}

// ClientLogLevel sets log level for the client, defaults to info
type ClientLogLevel int8

// A Config provides client configuration for sending requests to Twitter API
type Config struct {
	// The credentials to use when signing requests
	Credentials *Credentials

	// The HTTP client to use when sending requests. Defaults to `http.DefaultClient`.
	HTTPClient *http.Client

	// The logger writer interface to write logging messages to. Defaults to standard out.
	Logger *zerolog.Logger

	// The maximum number of times that a request will be retried for failures.
	MaxRetries int

	// Retryer guides how HTTP requests should be retried in case of
	// recoverable failures.
	//
	// When nil or the value does not implement the request.Retryer interface,
	// the client.DefaultRetryer will be used.
	Retryer RequestRetryer

	// See the ClientLogMode type documentation for the complete set of logging modes and available
	// configuration.
	ClientLogLevel ClientLogLevel
}

// NewConfig returns a new Config pointer that can be chained with builder
// methods to set multiple configuration values inline without using pointers.
func NewConfig() *Config {
	return &Config{}
}

// WithCredentials sets a config Credentials returning a Config pointer for chaining.
func (c *Config) WithCredentials(creds *Credentials) *Config {
	c.Credentials = creds
	return c
}

// WithHTTPClient sets a config HTTPClient value returning a Config pointe for chaining.
func (c *Config) WithHTTPClient(client *http.Client) *Config {
	c.HTTPClient = client
	return c
}

// WithLogger sets a config Logger value returning a Config pointer for chaining.
func (c *Config) WithLogger(logger zerolog.Logger) *Config {
	c.Logger = &logger
	return c
}

// WithLogLevel sets the client logging level
func (c *Config) WithLogLevel(logLevel int8) *Config {
	if c.Logger != nil {
		c.Logger.Level(zerolog.Level(logLevel))
	}
	return c
}

// WithMaxRetries sets a config MaxRetries value returning a Config pointer for chaining.
func (c *Config) WithMaxRetries(max int) *Config {
	c.MaxRetries = max
	return c
}

// WithRetryer sets a config Retryer value returning a Config pointer for chaining.
func (c *Config) WithRetryer(retryer RequestRetryer) *Config {
	c.Retryer = retryer
	return c
}

// NewDefaultLogger returns a Logger which will write log messages to stdout.
func newDefaultLogger() zerolog.Logger {
	return zerolog.New(os.Stderr).With().Timestamp().Logger()
}

func resolveConfig(cfg *Config) *Config {
	if cfg.HTTPClient == nil {
		cfg = cfg.WithHTTPClient(&http.Client{})
	}
	if cfg.Logger == nil {
		cfg = cfg.WithLogger(newDefaultLogger())
	}
	if cfg.MaxRetries == 0 {
		cfg = cfg.WithMaxRetries(maxDefaultRetries)
	}

	return cfg
}
