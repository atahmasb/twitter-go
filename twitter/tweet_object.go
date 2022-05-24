package twitter

import (
	"time"
)

// Tweet is the basic building block of Twitter.
type Tweet struct {
	// The unique identifier of Tweet.
	ID string `json:"id,required"`

	// The actual UTF-8 text of the Tweet.
	Text string `json:"text,required"`

	// Specifies the type of attachments (if any) present in this Tweet
	Attachments Attachment `json:"attachments,omitempty"`

	// The unique identifier of the User who posted this Tweet.
	AuthorID string `json:"author_id,omitempty"`

	// Contains context annotations for the Tweet.
	ContextAnnotations []ContextAnnotation `json:"context_annotations,omitempty"`

	// The Tweet ID of the original Tweet of the conversation.
	ConversionID string `json:"conversation_id,omitempty"`

	// Creation time of the Tweet.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Entities which have been parsed out of the text of the Tweet.
	Entities TweetEntities `json:"entities,omitempty"`

	// GEO contains details about the location tagged by the user in Tweet.
	Geo Geo `json:"geo,omitempty"`

	// 	If the represented Tweet is a reply, this field will contain the original Tweet’s author ID.
	// This will not necessarily always be the user directly mentioned in the Tweet.
	InReplyToUserID string `json:"in_reply_to_user_id,omitempty"`

	// Language of Tweet.
	Lang string `json:"lang,omitempty"`

	// Engagement metrics for Tweet at the time of request.
	NonPublicMetrics TweetNonPublicMetrics `json:"non_public_metrics,omitempty"`

	// Engagement metrics, tracked in an organic context, for Tweet at the time of the request.
	OrganicMetrics TweetOrganicMetrics `json:"organic_metrics,omitempty"`

	// This field only surfaces when a Tweet contains a link.
	// The meaning of the field doesn’t pertain to the Tweet content itself,
	// but instead it is an indicator that the URL contained in the Tweet may
	// contain content or media identified as sensitive content.
	PossibleSensitive bool `json:"possibly_sensitive,omitempty"`

	// Engagement metrics, tracked in a promoted context, for Tweet at the time of the request.
	PromotedMetrics TweetPromotedMetrics `json:"promoted_metrics,omitempty"`

	// Public engagement metrics for Tweet at the time of the request.
	PublicMetrics TweetPublicMetrics `json:"public_metrics,omitempty"`

	// A list of Tweets this Tweet refers to.
	ReferencedTweets []ReferencedTweet `json:"referenced_tweets,omitempty"`

	// 	Shows you who can reply to a given Tweet.
	// Fields returned are "everyone", "mentioned_users", and "followers".
	ReplySettings string `json:"reply_settings,omitempty"`

	// Determine if a Twitter user posted from the web, mobile device, or other app.
	Source string `json:"source,omitempty"`

	// When present, contains withholding details for withheld content.
	// https://help.twitter.com/en/rules-and-policies/tweet-withheld-by-country
	Withheld Withheld `json:"withheld,omitempty"`
}

// Attachment is the type of attachments present in Tweet.
type Attachment struct {
	PollIDs   []string `json:"poll_ids"`
	MediaKeys []string `json:"media_keys"`
}

// ContextAnnotation contains context annotations for Tweet.
type ContextAnnotation struct {
	Domain struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"domain"`

	Entity struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"entity"`
}

// TweetEntities are extarc information parsed out of the text of Tweet.
type TweetEntities struct {
	Annotations []Annotation `json:"annotations"`
	Cashtags    []Cashtag    `json:"cashtags"`
	Hashtags    []Hashtag    `json:"hashtag"`
	Mentions    []Mention    `json:"mention"`
	URLs        []URL        `json:"urls"`
}

// Annotation is annotation found in the text of Tweet.
type Annotation struct {
	Start          int     `json:"start"`
	End            int     `json:"end"`
	Probability    float64 `json:"probability"`
	Type           string  `json:"type"`
	NormalizedText string  `json:"normalized_text"`
}

// Geo Contains details about the location tagged by the user in Tweet.
type Geo struct {
	Coordinates struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"coordinates"`
	PlaceID string `json:"place_id"`
}

// TweetNonPublicMetrics contains engagement metrics for Tweet.
type TweetNonPublicMetrics struct {
	ImpressionCount   int `json:"impression_count"`
	URLLinkClicks     int `json:"url_link_clicks"`
	UserProfileClicks int `json:"user_profile_clicks"`
}

// TweetOrganicMetrics contains engagement metrics for Tweet tracked in organic context.
type TweetOrganicMetrics struct {
	ImpressionCount   int `json:"impression_count"`
	URLLinkClicks     int `json:"url_link_clicks"`
	LikeCount         int `json:"like_count"`
	UserProfileClicks int `json:"user_profile_clicks"`
	ReplyCount        int `json:"reply_click"`
	RetweetCount      int `json:"retweet_count"`
}

// TweetPromotedMetrics contains engagement metrics for Tweet tracked in promoted context.
type TweetPromotedMetrics struct {
	ImpressionCount   int `json:"impression_count"`
	URLLinkClicks     int `json:"url_link_clicks"`
	LikeCount         int `json:"like_count"`
	UserProfileClicks int `json:"user_profile_clicks"`
	ReplyCount        int `json:"reply_click"`
	RetweetCount      int `json:"retweet_count"`
}

// TweetPublicMetrics contains engagement metrics for public Tweet.
type TweetPublicMetrics struct {
	QuoteCount   int `json:"quote_count"`
	LikeCount    int `json:"like_count"`
	ReplyCount   int `json:"reply_click"`
	RetweetCount int `json:"retweet_count"`
}

// ReferencedTweet contains id and type of  a referenced tweet.
type ReferencedTweet struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}
