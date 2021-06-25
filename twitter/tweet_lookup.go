package twitter

const (
	lookupTweets = "lookupTweets"
)

// LookupTweetsInput contains input data to include in query parameters in the request
type LookupTweetsInput struct {
	Fields QueryParamsFields
	IDs    []string
}

// LookupTweetsOutput contains tweets lookup endpoint output
type LookupTweetsOutput struct {
	Data     []Tweet      `json:"data"`
	Includes Includes     `json:"includes"`
	Errors   []Diagnostic `json:"errors"`
}

// LookupSingleTweetInput contains input data
type LookupSingleTweetInput struct {
	ID string `json:"id"`
}

// LookupSingleTweetOutput contains single tweet lookup endpoint output
type LookupSingleTweetOutput struct {
	Data     Tweet        `json:"data"`
	Includes Includes     `json:"includes"`
	Errors   []Diagnostic `json:"errors"`
}

// LookupTweets returns a variety of information about the Tweet
// specified by the requested ID or list of IDs.
func (c *Client) LookupTweets(input *LookupTweetsInput) (req *Request, output *LookupTweetsOutput) {
	queryParams := getQueryParamsFromTweetsInput(input.Fields)
	endpoint := &EndPointInfo{
		Name:        lookupTweets,
		HTTPMethod:  "GET",
		HTTPPath:    "tweets",
		QueryParams: queryParams,
	}

	output = &LookupTweetsOutput{}
	req = c.NewRequest(endpoint, nil, output)
	return
}

// LookupSingleTweet returns a variety of information about a single
// Tweet specified by the requested ID.
func (c *Client) LookupSingleTweet(input *LookupTweetsInput) (req *Request, output *LookupSingleTweetOutput) {

	endpoint := &EndPointInfo{
		Name:       lookupTweets,
		HTTPMethod: "GET",
		HTTPPath:   "tweets",
	}

	output = &LookupSingleTweetOutput{}
	req = c.NewRequest(endpoint, input, output)
	return
}
