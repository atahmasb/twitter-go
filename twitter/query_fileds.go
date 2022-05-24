package twitter

// ExpansionField enables you to select which specific expansion
// fields will deliver in each returned Tweet.
type Field string

func (e Field) Stringify() string {
	return string(e)
}

const (
	// ExpansionAttachmentsPollIDs returns a poll object containing metadata for the poll included in the Tweet
	ExpansionAttachmentsPollIDs Field = "attachments.poll_ids"
	// ExpansionAttachmentsMediaKeys returns a media object representing the images, videos, GIFs included in the Tweet
	ExpansionAttachmentsMediaKeys Field = "attachments.media_keys"
	// ExpansionAuthorID returns a user object representing the Tweet’s author
	ExpansionAuthorID Field = "author_id"
	// ExpansionEntitiesMentionsUserName returns a user object for the user mentioned in the Tweet
	ExpansionEntitiesMentionsUserName Field = "entities.mentions.username"
	// ExpansionGeoPlaceID returns a place object containing metadata for the location tagged in the Tweet
	ExpansionGeoPlaceID Field = "geo.place_id"
	// ExpansionInReplyToUserID returns a user object representing the Tweet author this requested Tweet is a reply of
	ExpansionInReplyToUserID Field = "in_reply_to_user_id"
	// ExpansionReferencedTweetsID returns a Tweet object that this Tweet is referencing (either as a Retweet, Quoted Tweet, or reply)
	ExpansionReferencedTweetsID Field = "referenced_tweets.id"
	// ExpansionReferencedTweetsIDAuthorID returns a user object for the author of the referenced Tweet
	ExpansionReferencedTweetsIDAuthorID Field = "referenced_tweets.id.author_id"
	// ExpansionPinnedTweetID returns a Tweet object representing the Tweet pinned to the top of the user’s profile
	ExpansionPinnedTweetID Field = "pinned_tweet_id"
)

const (
	// MediaFieldDurationMS available when type is video. Duration in milliseconds of the video.
	MediaFieldDurationMS Field = "duration_ms"
	// MediaFieldHeight of this content in pixels.
	MediaFieldHeight Field = "height"
	// MediaFieldMediaKey unique identifier of the expanded media content.
	MediaFieldMediaKey Field = "media_key"
	// MediaFieldPreviewImageURL is the URL to the static placeholder preview of this content.
	MediaFieldPreviewImageURL Field = "preview_image_url"
	// MediaFieldType is the type of content (animated_gif, photo, video)
	MediaFieldType Field = "type"
	// MediaFieldURL is the URL of the content
	MediaFieldURL Field = "url"
	// MediaFieldWidth is the width of this content in pixels
	MediaFieldWidth Field = "width"
	// MediaFieldPublicMetrics is the public engagement metrics for the media content at the time of the request.
	MediaFieldPublicMetrics Field = "public_metrics"
	// MediaFieldNonPublicMetrics is the non-public engagement metrics for the media content at the time of the request.
	MediaFieldNonPublicMetrics Field = "non_public_metrics"
	// MediaFieldOrganicMetrics is the engagement metrics for the media content, tracked in an organic context, at the time of the request.
	MediaFieldOrganicMetrics Field = "organic_metrics"
	// MediaFieldPromotedMetrics is the URL to the static placeholder preview of this content.
	MediaFieldPromotedMetrics Field = "promoted_metrics"
)

const (
	// PlaceFieldContainedWithin returns the identifiers of known places that contain the referenced place.
	PlaceFieldContainedWithin Field = "contained_within"
	// PlaceFieldCountry is the full-length name of the country this place belongs to.
	PlaceFieldCountry Field = "country"
	// PlaceFieldCountryCode is the ISO Alpha-2 country code this place belongs to.
	PlaceFieldCountryCode Field = "country_code"
	// PlaceFieldFullName is a longer-form detailed place name.
	PlaceFieldFullName Field = "full_name"
	// PlaceFieldGeo contains place details in GeoJSON format.
	PlaceFieldGeo Field = "geo"
	// PlaceFieldID is the unique identifier of the expanded place, if this is a point of interest tagged in the Tweet.
	PlaceFieldID Field = "id"
	// PlaceFieldName is the short name of this place
	PlaceFieldName Field = "name"
	// PlaceFieldPlaceType is specified the particular type of information represented by this place information, such as a city name, or a point of interest.
	PlaceFieldPlaceType Field = "place_type"
)

const (
	// PollFieldDurationMinutes specifies the total duration of this poll.
	PollFieldDurationMinutes Field = "duration_minutes"
	// PollFieldEndDateTime specifies the end date and time for this poll.
	PollFieldEndDateTime Field = "end_datetime"
	// PollFieldID is unique identifier of the expanded poll.
	PollFieldID Field = "id"
	// PollFieldOptions contains objects describing each choice in the referenced poll.
	PollFieldOptions Field = "options"
	// PollFieldVotingStatus indicates if this poll is still active and can receive votes, or if the voting is now closed.
	PollFieldVotingStatus Field = "voting_status"
)

const (
	// TweetFieldID is the unique identifier of the requested Tweet.
	TweetFieldID Field = "id"
	// TweetFieldText is the actual UTF-8 text of the Tweet. See twitter-text for details on what characters are currently considered valid.
	TweetFieldText Field = "text"
	// TweetFieldAttachments specifies the type of attachments (if any) present in this Tweet.
	TweetFieldAttachments Field = "attachments"
	// TweetFieldAuthorID is the unique identifier of the User who posted this Tweet
	TweetFieldAuthorID Field = "author_id"
	// TweetFieldContextAnnotations contains context annotations for the Tweet.
	TweetFieldContextAnnotations Field = "context_annotations"
	// TweetFieldConversationID is the Tweet ID of the original Tweet of the conversation (which includes direct replies, replies of replies).
	TweetFieldConversationID Field = "conversation_id"
	// TweetFieldCreatedAt is the creation time of the Tweet.
	TweetFieldCreatedAt Field = "created_at"
	// TweetFieldEntities are the entities which have been parsed out of the text of the Tweet. Additionally see entities in Twitter Objects.
	TweetFieldEntities Field = "entities"
	// TweetFieldGeo contains details about the location tagged by the user in this Tweet, if they specified one.
	TweetFieldGeo Field = "geo"
	// TweetFieldInReplyToUserID will contain the original Tweet’s author ID
	TweetFieldInReplyToUserID Field = "in_reply_to_user_id"
	// TweetFieldLanguage is the language of the Tweet, if detected by Twitter. Returned as a BCP47 language tag.
	TweetFieldLanguage Field = "lang"
	// TweetFieldNonPublicMetrics are the non-public engagement metrics for the Tweet at the time of the request.
	TweetFieldNonPublicMetrics Field = "non_public_metrics"
	// TweetFieldPublicMetrics are the public engagement metrics for the Tweet at the time of the request.
	TweetFieldPublicMetrics Field = "public_metrics"
	// TweetFieldOrganicMetrics are the engagement metrics, tracked in an organic context, for the Tweet at the time of the request.
	TweetFieldOrganicMetrics Field = "organic_metrics"
	// TweetFieldPromotedMetrics are the engagement metrics, tracked in a promoted context, for the Tweet at the time of the request.
	TweetFieldPromotedMetrics Field = "promoted_metrics"
	// TweetFieldPossiblySensitive is an indicator that the URL contained in the Tweet may contain content or media identified as sensitive content.
	TweetFieldPossiblySensitive Field = "possibly_sensitive"
	// TweetFieldReferencedTweets is a list of Tweets this Tweet refers to.
	TweetFieldReferencedTweets Field = "referenced_tweets"
	// TweetFieldSource is the name Field the app the user Tweeted from.
	TweetFieldSource Field = "source"
	// TweetFieldWithHeld contains withholding details
	TweetFieldWithHeld Field = "withheld"
)

const (
	// UserFieldCreatedAt is the UTC datetime that the user account was created on Twitter.
	UserFieldCreatedAt Field = "created_at"
	// UserFieldDescription is the text of this user's profile description (also known as bio), if the user provided one.
	UserFieldDescription Field = "description"
	// UserFieldEntities contains details about text that has a special meaning in the user's description.
	UserFieldEntities Field = "entities"
	// UserFieldID is the unique identifier of this user.
	UserFieldID Field = "id"
	// UserFieldLocation is the location specified in the user's profile, if the user provided one.
	UserFieldLocation Field = "location"
	// UserFieldName is the name of the user, as they’ve defined it on their profile
	UserFieldName Field = "name"
	// UserFieldPinnedTweetID is the unique identifier of this user's pinned Tweet.
	UserFieldPinnedTweetID Field = "pinned_tweet_id"
	// UserFieldProfileImageURL is the URL to the profile image for this user, as shown on the user's profile.
	UserFieldProfileImageURL Field = "profile_image_url"
	// UserFieldProtected indicates if this user has chosen to protect their Tweets (in other words, if this user's Tweets are private).
	UserFieldProtected Field = "protected"
	// UserFieldPublicMetrics contains details about activity for this user.
	UserFieldPublicMetrics Field = "public_metrics"
	// UserFieldURL is the URL specified in the user's profile, if present.
	UserFieldURL Field = "url"
	// UserFieldUserName is the Twitter screen name, handle, or alias that this user identifies themselves with
	UserFieldUserName Field = "username"
	// UserFieldVerified indicates if this user is a verified Twitter User.
	UserFieldVerified Field = "verified"
	// UserFieldWithHeld contains withholding details
	UserFieldWithHeld Field = "withheld"
)
