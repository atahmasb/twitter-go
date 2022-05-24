package twitter

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"
)

const (
	validateRules = "validateRules"
	createRules   = "createRules"
	deleteRules   = "deleteRules"
	getRules      = "getRules"
	streamTweets  = "streamTweets"
)

// Rule contains a rule's value, tag, and ID
type Rule struct {
	Value string      `json:"value"`
	Tag   string      `json:"tag"`
	ID    json.Number `json:"id"`
}

// ValidateRulesInput is used to validate a set of rules syntax before creating the rule
type ValidateRulesInput struct {
	Add []Rule `json:"add,required"`
}

// ValidateRulesOutputMetaSummary contains information about number of corret and incorrect rules
type ValidateRulesOutputMetaSummary struct {
	Created    int `json:"created"`
	NotCreated int `json:"not_created"`
}

// ValidateRulesOutputMeta contains meta information about validating rules endpoint response
type ValidateRulesOutputMeta struct {
	Sent    time.Time                      `json:"sent"`
	Summary ValidateRulesOutputMetaSummary `json:"summary"`
}

// ValidateRulesOutput contains output of validating rules endpoint response
type ValidateRulesOutput struct {
	Data []Rule                  `json:"data"`
	Meta ValidateRulesOutputMeta `json:"meta"`
}

// CreateRulesInput is used to create a set of rules
type CreateRulesInput struct {
	Add []Rule `json:"add,required"`
}

// CreateRulesOutput contains output of creating rules endpoint response
type CreateRulesOutput struct {
	Data []Rule                  `json:"data"`
	Meta ValidateRulesOutputMeta `json:"meta"`
}

// RulesIDs contains a list of rule ids to be deleted
type RulesIDs struct {
	IDs []string `json:"ids"`
}

// DeleteRulesInput contains input request to delete rules endpoint
type DeleteRulesInput struct {
	Delete RulesIDs `json:"delete"`
}

// DeleteRulesOuput contains output of deleting rules endpoint response
type DeleteRulesOuput struct {
	Meta DeleteRulesOutputMeta `json:"meta"`
}

// DeleteRulesOutputMeta contains meta information about deleting rules endpoint response
type DeleteRulesOutputMeta struct {
	Sent    time.Time                    `json:"sent"`
	Summary DeleteRulesOutputMetaSummary `json:"summary"`
}

// DeleteRulesOutputMetaSummary contains information about number of deleted and not deleted rules
type DeleteRulesOutputMetaSummary struct {
	Deleted    int `json:"deleted"`
	NotDeleted int `json:"not_deleted"`
}

// GetRulesInput contains input request to retrieving rules endpoint
type GetRulesInput struct {
	IDs []string `json:"ids"`
}

// GetRulesOutput contains output of retrieving rules endpoint response
type GetRulesOutput struct {
	Data []Rule `json:"data"`
}

// StreamTweetsInput contains input query parameters to include in the request
type StreamTweetsInput struct {
	Expansions []Field
	Media      []Field
	Place      []Field
	Poll       []Field
	Tweet      []Field
	User       []Field
}

// StreamTweetsOutput contains streaming endpoint output
type StreamTweetsOutput struct {
	Data          Tweet        `json:"data"`
	Includes      Includes     `json:"includes"`
	Errors        []Diagnostic `json:"errors"`
	MatchingRules []Rule       `json:"matching_rules"`
}

// ValidateRules tests the syntax of your rule without submitting it
func (c *Client) ValidateRules(input *ValidateRulesInput) (req *Request, output *ValidateRulesOutput) {
	queryParams := make(map[string]string)
	queryParams["dry_run"] = "true"

	endpoint := &EndPointInfo{
		Name:        validateRules,
		HTTPMethod:  "POST",
		HTTPPath:    "tweets/search/stream/rules",
		QueryParams: queryParams,
	}

	if input == nil {
		input = &ValidateRulesInput{}
	}

	output = &ValidateRulesOutput{}
	req = c.NewRequest(endpoint, input, output)
	return
}

// CreateRules adds rules to your stream
func (c *Client) CreateRules(input *CreateRulesInput) (req *Request, output *CreateRulesOutput) {
	endpoint := &EndPointInfo{
		Name:       createRules,
		HTTPMethod: "POST",
		HTTPPath:   "tweets/search/stream/rules",
	}

	if input == nil {
		input = &CreateRulesInput{}

	}

	output = &CreateRulesOutput{}
	req = c.NewRequest(endpoint, input, output)
	return
}

// DeleteRules removes rules from your stream
func (c *Client) DeleteRules(input *DeleteRulesInput) (req *Request, output *DeleteRulesOuput) {
	endpoint := &EndPointInfo{
		Name:       deleteRules,
		HTTPMethod: "POST",
		HTTPPath:   "tweets/search/stream/rules",
	}

	if input == nil {
		input = &DeleteRulesInput{}
	}

	output = &DeleteRulesOuput{}
	req = c.NewRequest(endpoint, input, output)
	return
}

// GetRules retrives rules that have been applied to your stream
func (c *Client) GetRules(input *GetRulesInput) (req *Request, output *GetRulesOutput) {
	queryParams := make(map[string]string)
	if len(input.IDs) > 0 {
		queryParams["ids"] = strings.Join(input.IDs, ",")
	}

	endpoint := &EndPointInfo{
		Name:        getRules,
		HTTPMethod:  "GET",
		HTTPPath:    "tweets/search/stream/rules",
		QueryParams: queryParams,
	}

	if input == nil {
		input = &GetRulesInput{}
	}

	output = &GetRulesOutput{}
	req = c.NewRequest(endpoint, input, output)
	return

}

// StreamTweets streams Tweets in real-time based on a specific set of filter rules.
// You need to check if stream is established by checking `receiving` parameter of the
// Stream struct. Streaming tweets can be accessed through the Queue on the return
// stream struct
func (c *Client) StreamTweets(input StreamTweetsInput) (s *Stream) {
	queryParams := getQueryParamsFromStreamTweetsInput(input)
	endpoint := &EndPointInfo{
		Name:        streamTweets,
		HTTPMethod:  "GET",
		HTTPPath:    "tweets/search/stream",
		QueryParams: queryParams,
	}

	output := &StreamTweetsOutput{}
	s = c.NewStream(endpoint, nil, output)
	return

}

func getQueryParamsFromStreamTweetsInput(input StreamTweetsInput) map[string]string {
	queryParams := make(map[string]string, 0)
	fields := reflect.Indirect(reflect.ValueOf(&input))
	numberOfFields := fields.NumField()
	for i := 0; i < numberOfFields; i++ {
		fieldName := fields.Type().Field(i).Name
		fieldValue := fields.Field(i).Interface()
		fieldParams, ok := fieldValue.([]Field)
		if !ok {
			continue
		}
		if len(fieldParams) == 0 {
			continue
		}

		switch fieldName {
		case "Expansions":
			queryParams["expansions"] = joinFieldParams(fieldParams)
		case "Media":
			queryParams["media.fields"] = joinFieldParams(fieldParams)
		case "Place":
			queryParams["place.fields"] = joinFieldParams(fieldParams)
		case "Poll":
			queryParams["poll.fields"] = joinFieldParams(fieldParams)
		case "Tweet":
			queryParams["tweet.fields"] = joinFieldParams(fieldParams)
		case "User":
			queryParams["user.fields"] = joinFieldParams(fieldParams)
		default:
			continue

		}

	}
	return queryParams
}

func joinFieldParams(params []Field) string {
	fieldsAll := make([]string, len(params), len(params))
	for idx, field := range params {
		fieldsAll[idx] = field.Stringify()
	}
	return strings.Join(fieldsAll, ",")

}
