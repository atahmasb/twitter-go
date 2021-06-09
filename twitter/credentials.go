package twitter

import (
	"errors"
)

var (
	// ErrCredentialsEmpty is emitted when credentials are empty.
	ErrCredentialsEmpty = errors.New("EmptyCredentials: Credentials value is empty")
)

// Value is the Twitter credentials values.
type Value struct {
	BearerToken string
}

// A Credentials is a set of credentials which are set programmatically,
type Credentials struct {
	Value
}

// NewCredentials returns a pointer to a new Credentials with the provider set.
func NewCredentials(creds Value) *Credentials {
	c := &Credentials{
		Value: creds,
	}
	return c
}

// Retrieve returns the credentials or error if the credentials are invalid.
func (c *Credentials) Retrieve() (Value, error) {
	if c.BearerToken == "" {
		return Value{}, ErrCredentialsEmpty
	}
	return c.Value, nil
}

// IsSet checks if the token is set on the credential struct
func (c *Credentials) IsSet() bool {
	if c.BearerToken == "" {
		return false
	}
	return true
}
