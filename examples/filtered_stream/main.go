package main

import (

	"os"
	"time"

	"github.com/atahmasb/twitter-go/examples"
	"github.com/atahmasb/twitter-go/twitter"
	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	cred := twitter.NewCredentials(twitter.Value{BearerToken: os.Getenv("BEARER_TOKEN")})
	cfg := twitter.NewConfig().WithCredentials(cred)
	client := twitter.NewClient(cfg)

	// Set a rule value and tag
	// Visit https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/integrate/build-a-rule
	// To lean about building rules and different operations available
	value := "Example of using Twitter filtered stream endpoints"
	tag := "Twitter API v2"

	// Validate rules
	logger.Info().Msg("Validating rule")
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
		logger.Error().Err(validateRuleErr).Msg("Failed to validate rule")
		os.Exit(1)
	}

	validateRuleOutputMap, err := examples.StructToMap(validateRuleOutput)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to convert validate rule output to struct to map")
	}
	logger.Info().Msgf("Validate rule's output: %v", validateRuleOutputMap)

	// Create rules
	logger.Info().Msg("Creating rule")
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
		logger.Error().Err(createRuleErr).Msg("Failed to create rule")
		os.Exit(1)
	}

	createRuleOutputMap, err := examples.StructToMap(createRuleOutput)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to convert create rule output struct to map")
	}
	logger.Info().Msgf("Create rule's output: %v", createRuleOutputMap)

	// Streams Tweets in real-time based on a specific set of filter rules
	logger.Info().Msgf("Streaming real time tweets")
	stream := client.StreamTweets(twitter.StreamTweetsInput{})

	// Stream tweets for 5 seconds and then call`Stop` to stop streaming
	go func() {
		time.Sleep(time.Second * 5)
		logger.Info().Msg("Stop streaming real time tweets")
		stream.Stop()

	}()
	for message := range stream.MessageQueue {

		streamOutput, ok := message.(*twitter.StreamTweetsOutput)
		if !ok {
			logger.Error().Msg("Failed to cast message to Tweet")
		} else {
			logger.Info().Msgf("Tweet ID: %s, text: %s", streamOutput.Data.ID, streamOutput.Data.Text)
		}

	}

	// Get rules
	getRuleReq, getRuleOutput := client.GetRules(&twitter.GetRulesInput{})

	getRuleErr := getRuleReq.Send()
	if getRuleErr != nil {
		logger.Error().Err(getRuleErr).Msg("Failed to get rule")
		os.Exit(1)
	}

	getRuleOutputMap, err := examples.StructToMap(getRuleOutput)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to convert get rule output struct to map")
	}
	logger.Info().Msgf("Get rule's output: %v", getRuleOutputMap)

	// Delete rules
	if len(createRuleOutput.Data) > 0 {
		logger.Info().Msg("Deleting rule")
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
			logger.Error().Err(deleteRuleErr).Msg("Failed to delete rule")
			os.Exit(1)
		}

		deleteRuleOutputMap, err := examples.StructToMap(deleteRuleOutput)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to convert delete rule output struct to map")
		}
		logger.Info().Msgf("Delete rule's output: %v", deleteRuleOutputMap)
	}

}
