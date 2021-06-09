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
