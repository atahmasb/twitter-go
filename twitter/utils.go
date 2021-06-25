package twitter

import (
	"reflect"
	"strings"
)

// Field enables you to select which specific expansion
// fields will deliver in each returned Tweet.
type Field string

func (e Field) stringify() string {
	return string(e)
}

// QueryParamsFields contains query parameters options that if passed
// it's data will be returned in the response tweets
type QueryParamsFields struct {
	Expansions []Field
	Media      []Field
	Place      []Field
	Poll       []Field
	Tweet      []Field
	User       []Field
}

func getQueryParamsFromTweetsInput(input QueryParamsFields) map[string]string {
	queryParams := make(map[string]string, 0)
	fields := reflect.Indirect(reflect.ValueOf(input))
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
		fieldsAll[idx] = field.stringify()
	}
	return strings.Join(fieldsAll, ",")

}
