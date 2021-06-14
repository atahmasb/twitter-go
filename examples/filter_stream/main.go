package main

import (
	"fmt"
	"os"
	"time"

	"github.com/atahmasb/twitter-go/twitter"
)

func main() {
	cred := twitter.NewCredentials(twitter.Value{BearerToken: os.Getenv("BEARER_TOKEN")})
	cfg := twitter.NewConfig().WithCredentials(cred)
	client := twitter.NewClient(cfg)

	// Set a rule value and tag
	// Visit https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/integrate/build-a-rule
	// To lean about building rules and different operations available
	value := "Example of using Twitter filter stream endpoints"
	tag := "Twitter API v2"

	// Validate a new rule
	fmt.Println("Validating rule")
	validateRuleInput := twitter.ValidateRulesInput{
		Add: []twitter.Rule{
			{
				Value: value,
				Tag:   tag,
			},
		},
	}

	validateRuleReq, validateRuleOutput := client.ValidateRules(&validateRuleInput)

	validateRuleErr := validateRuleReq.Send()
	if validateRuleErr != nil {
		fmt.Println(validateRuleErr)
		os.Exit(1)
	}

	fmt.Println(validateRuleOutput)

	// Create a new rule
	fmt.Println("Creating rule")
	createRuleinput := twitter.CreateRulesInput{
		Add: []twitter.Rule{
			{
				Value: value,
				Tag:   tag,
			},
		},
	}
	createRuleReq, createRuleOutput := client.CreateRules(&createRuleinput)

	createRuleErr := createRuleReq.Send()
	if createRuleErr != nil {
		fmt.Println(createRuleErr)
		os.Exit(1)
	}

	fmt.Println(createRuleOutput)

	// Streams Tweets in real-time based on a specific set of filter rules
	stream := client.StreamTweets(twitter.StreamTweetsInput{})

	// Stream tweets for 5 seconds and then call`Stop` to stop streaming
	fmt.Println("Streaming real time tweets")
	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("Stop streaming real time tweets")
		stream.Stop()

	}()
	for message := range stream.MessageQueue {

		streamOutput, ok := message.(*twitter.StreamTweetsOutput)
		if !ok {
			fmt.Println("Failed to cast message to Tweet")
		} else {
			fmt.Printf("tweet id: %s, tweet text: %s", streamOutput.Data.ID, streamOutput.Data.Text)
		}

	}

	// Delete rule
	fmt.Println("Deleting rule")
	if len(createRuleOutput.Data) > 0 {
		ruleID := createRuleOutput.Data[0].ID.String()

		deleteRuleInput := twitter.DeleteRulesInput{
			Delete: twitter.RulesIDs{
				IDs: []string{
					ruleID,
				},
			}}

		deleteRuleReq, deleteRuleOutput := client.DeleteRules(&deleteRuleInput)

		deleteRuleErr := deleteRuleReq.Send()
		if deleteRuleErr != nil {
			fmt.Println(deleteRuleErr)
			os.Exit(1)
		}

		fmt.Println(deleteRuleOutput)
	}

}
