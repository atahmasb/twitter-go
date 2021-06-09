package twitter

import (
	"time"
)

// Retryer provides the interface for the client request retry behavior. The
// Retryer implementation is responsible for implementing exponential backoff,
// and determine if a request API error should be retried.
type Retryer interface {
	// RetryRules return the retry delay that should be used by the SDK before
	// making another request attempt for the failed request.
	RetryRules(*Request) time.Duration

	// ShouldRetry returns if the failed request is retryable.
	//
	// Implementations may consider request attempt count when determining if a
	// request is retryable, but the client will use MaxRetries to limit the
	// number of attempts a request are made.
	ShouldRetry(*Request) bool

	// MaxRetries is the number of times a request may be retried before
	// failing.
	MaxRetries() int
}

// noRetryer should be used when a request is created without a retryer.
type noRetryer struct{}

func (d noRetryer) MaxRetries() int {
	return 0
}

func (d noRetryer) ShouldRetry(*Request) bool {
	return false
}

func (d noRetryer) RetryRules(*Request) time.Duration {
	return 0
}