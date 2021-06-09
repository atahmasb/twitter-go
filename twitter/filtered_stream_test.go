package twitter

import (
	"fmt"
	"net/http"
)

func (suite *twitterClientSuite) Test_ValidateRules() {
	expectedMethod := "POST"
	suite.mux.HandleFunc("/2/tweets/search/stream/rules", func(w http.ResponseWriter, r *http.Request) {
		suite.assertMethod(expectedMethod, r)
		suite.assertQuery(map[string]string{"dry_run": "true"}, r)
		w.Header().Set("Content/type", "application/json")
		fmt.Fprintf(w, `{"data": [{"value": "meme", "tag": "funny things", "id": "1166895166390583299"},{"value": "cats has:media -grumpy", "tag": "happy cats with media", "id": "1166895166390583296"}],"meta": {"sent": "2019-08-29T02:07:42.205Z", "summary": {"created": 2,"not_created": 0}}}`)
	})

	input := &ValidateRulesInput{
		[]Rule{{
			Value: "cats has:media -grumpy",
			Tag:   "happy cats with media",
		},
			{
				Value: "funny things",
				Tag:   "funny things",
			}},
	}

	req, out := suite.client.ValidateRules(input)
	err := req.Send()

	suite.Assert().Nil(err)
	suite.Assert().NotNil(out)
	suite.Assert().NotNil(out.Meta)
	suite.Assert().NotNil(out.Data)
	suite.Assert().Equal(2, len(out.Data))
}

func (suite *twitterClientSuite) Test_CreateRules() {
	expectedMethod := "POST"
	suite.mux.HandleFunc("/2/tweets/search/stream/rules", func(w http.ResponseWriter, r *http.Request) {
		suite.assertMethod(expectedMethod, r)
		w.Header().Set("Content/type", "application/json")
		fmt.Fprintf(w, `{"data": [{"value": "meme", "tag": "funny things", "id": "1166895166390583299"},{"value": "cats has:media -grumpy", "tag": "happy cats with media", "id": "1166895166390583296"}],"meta": {"sent": "2019-08-29T02:07:42.205Z", "summary": {"created": 2,"not_created": 0}}}`)
	})

	input := &CreateRulesInput{
		[]Rule{{
			Value: "cats has:media -grumpy",
			Tag:   "happy cats with media",
		},
			{
				Value: "funny things",
				Tag:   "funny things",
			}},
	}

	req, out := suite.client.CreateRules(input)
	err := req.Send()

	suite.Assert().Nil(err)
	suite.Assert().NotNil(out)
	suite.Assert().NotNil(out.Meta)
	suite.Assert().NotNil(out.Data)
	suite.Assert().Equal(2, len(out.Data))
}

func (suite *twitterClientSuite) Test_DeleteRules() {
	expectedMethod := "POST"
	suite.mux.HandleFunc("/2/tweets/search/stream/rules", func(w http.ResponseWriter, r *http.Request) {
		suite.assertMethod(expectedMethod, r)
		w.Header().Set("Content/type", "application/json")
		fmt.Fprintf(w, `{"meta": {"sent": "2019-08-29T01:48:54.633Z", "summary": {"deleted": 2, "not_deleted": 0}}}`)
	})

	input := &DeleteRulesInput{
		RulesIDs{
			[]string{
				"1166895166390583299",
				"1166895166390583296",
			},
		},
	}

	req, out := suite.client.DeleteRules(input)
	err := req.Send()

	suite.Assert().Nil(err)
	suite.Assert().NotNil(out)
	suite.Assert().NotNil(out.Meta)
	suite.Assert().Equal(2, out.Meta.Summary.Deleted)
	suite.Assert().Equal(0, out.Meta.Summary.NotDeleted)
}

func (suite *twitterClientSuite) Test_GetRules() {
	expectedMethod := "GET"
	suite.mux.HandleFunc("/2/tweets/search/stream/rules", func(w http.ResponseWriter, r *http.Request) {
		suite.assertMethod(expectedMethod, r)
		suite.assertQuery(map[string]string{"ids": "1165037377523306497,1165037377523306498"}, r)
		w.Header().Set("Content/type", "application/json")
		fmt.Fprintf(w, `{"data": [{"id": "1165037377523306497", "value": "dog has:images", "tag": "dog pictures"}, {"id": "1165037377523306498", "value": "cat has:images -grumpy"}], "meta": {"sent": "2019-08-29T01:12:10.729Z"}}`)})

	input := &GetRulesInput{
		[]string{
			"1165037377523306497",
			"1165037377523306498",
		},
	}

	req, out := suite.client.GetRules(input)
	err := req.Send()

	suite.Assert().Nil(err)
	suite.Assert().NotNil(out)
	suite.Assert().NotNil(out.Data)
	suite.Assert().Equal(2, len(out.Data))
}

func (suite *twitterClientSuite) Test_StreamTweets() {
		expectedMethod := "GET"
		suite.mux.HandleFunc("/2/tweets/search/stream", func(w http.ResponseWriter, r *http.Request) {
			suite.assertMethod(expectedMethod, r)
			w.Header().Set("Content/type", "application/json")
			fmt.Fprintf(w, `{"data": {"id": "1067094924124872705", "text": "Twitter API is awesome!"}}` + "\r\n" + `{"data": {"id": "1067094924124872705", "text": "Check this client out and connect to Twitter!"}}` + "\r\n")
		})

		input := StreamTweetsInput{}

		stream := suite.client.StreamTweets(input)
		
		for message := range stream.MessageQueue {
			suite.Assert().NotNil(message)
		}

		stream.Stop()
		suite.Assert().Equal(len(stream.MessageQueue), 0)
}
