package twitter

import (
	"time"
)

// User contains Twitter user account metadata describing user.
type User struct {
	// The unique identifier of User.
	ID string `json:"id,required"`

	// The name of the user, as they’ve defined it on their profile.
	// Not necessarily a person’s name. Typically capped at 50 characters,
	// but subject to change.
	Name string `json:"name,required"`

	// The Twitter screen name, handle, or alias that this user identifies
	// themselves with. Usernames are unique but subject to change. Typically
	// a maximum of 15 characters long, but some historical accounts may exist
	// with longer names.
	UserName string `json:"username,required"`

	// The UTC datetime that the user account was created on Twitter.
	CreatedAt time.Time `json:"created_at"`

	// The text of this user's profile description (also known as bio),
	// if the user provided one.
	Description string `json:"description"`

	// Entities which have been parsed out of the text of the Tweet.
	Entities UserEntities `json:"entities"`

	// The location specified in the user's profile, if the user provided one.
	// As this is a freeform value, it may not indicate a valid location,
	// but it may be fuzzily evaluated when performing searches with location queries.
	Location string `json:"location"`

	// Unique identifier of this user's pinned Tweet.
	PinnedTweetID string `json:"pinned_tweet_id"`

	// The URL to the profile image for this user, as shown on the user's profile.
	ProfileImageURL string `json:"profile_image_url"`

	// Indicates if this user has chosen to protect their Tweets
	// (in other words, if this user's Tweets are private).
	Protected bool `json:"protected"`

	// Contains details about activity for user.
	PublicMetrics UserPublicMetrics `json:"public_metrics"`

	// The URL specified in the user's profile, if present.
	URL string `json:"url"`

	// Indicates if this user is a verified Twitter User.
	Verified bool `json:"verified"`

	Withheld Withheld `json:"withheld"`
}

// UserEntities Contains details about text that has a special meaning in the user's description.
type UserEntities struct {
	Cashtags []Cashtag `json:"cashtags"`
	Hashtags []Hashtag `json:"hashtag"`
	Mentions []Mention `json:"mention"`
	URLs     []URL     `json:"urls"`
}

// UserPublicMetrics Contains details about activity for this user.
type UserPublicMetrics struct {
	FollowingCount int `json:"following_count"`
	FollowersCount int `json:"followers_count"`
	TweetCount     int `json:"tweet_count"`
	ListedCount    int `json:"listed_count"`
}
