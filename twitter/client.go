package twitter

// Twitter API information constants
const (
	EndPoint   = "https://api.twitter.com" // Base endpoint for for Twitter URL
	APIVersion = "2"                       // Twitter API version.
)

// APIInfo wraps immutable data
type APIInfo struct {
	Endpoint   string
	APIVersion string
}

// Client is twitter client to interact with Twitter API.
type Client struct {
	Config  *Config
	APIInfo APIInfo
	Retryer Retryer
}

// NewClient returns a new Twitter API client that uses default handlers and configs.
func NewClient(cfg *Config) *Client {
	cfg = resolveConfig(cfg)
	client := &Client{
		Config:  cfg,
		APIInfo: newAPIInfo(),
	}

	switch retryer, ok := cfg.Retryer.(Retryer); {
	case ok:
		client.Retryer = retryer
	case cfg.Retryer != nil && cfg.Logger != nil:
		cfg.Logger.Warn().Msgf("%T does not implement Retryer; using DefaultRetryer instead", cfg.Retryer)
		fallthrough
	default:
		cfg.Retryer = noRetryer{}
	}

	return client
}

// newAPIInfo returns twitter API basic information such as api endpoint and version
func newAPIInfo() APIInfo {
	return APIInfo{
		Endpoint:   EndPoint,
		APIVersion: APIVersion,
	}
}
