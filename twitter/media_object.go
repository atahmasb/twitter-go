package twitter

// Media refers to any image, GIF, or video attached to a Tweet.
type Media struct {

	// Unique identifier of the expanded media content.
	MediaKey string `json:"media_key,required"`

	// Type of content (animated_gif, photo, video).
	Type string `json:"type"`

	// Available when type is video. Duration in milliseconds of the video.
	Duration int `json:"duration_ms"`

	// Height of content in pixels.
	Height int `json:"height"`

	// Non-public engagement metrics for the media content at the time of the request.
	// Requires user context authentication.
	NonPublicMetrics MediaNonPublicMetrics `json:"non_public_metrics"`

	// Engagement metrics for the media content, tracked in an organic context,
	// at the time of the request. Requires user context authentication.
	OrganicMetrics MediaOrganicMetrics `json:"organic_metrics"`

	// URL to the static placeholder preview of this content.
	PreviewImageURL string `json:"preview_image_url"`

	// Engagement metrics for the media content, tracked in a promoted context,
	// at the time of the request. Requires user context authentication.
	PromotedMetrics MediaPromotedMetrics `json:"promoted_metrics"`

	// Public engagement metrics for the media content at the time of the request.
	PublicMetrics MediaPublicMetrics `json:"public_metrics"`

	// Width of this content in pixels.
	Width int `json:"width"`
}

// MediaNonPublicMetrics contains non-public engagement metrics for the media content at the time of the request.
type MediaNonPublicMetrics struct {
	PlayBack0Count   int `json:"playback_0_count"`
	PlayBack100Count int `json:"playback_100_count"`
	PlayBack25Count  int `json:"playback_25_count"`
	PlayBack50Count  int `json:"playback_50_count"`
	PlayBack75Count  int `json:"playback_75_count"`
}

// MediaOrganicMetrics contains engagement metrics for the media content, tracked in an organic context, at the time of the request.
type MediaOrganicMetrics struct {
	PlayBack0Count   int `json:"playback_0_count"`
	PlayBack100Count int `json:"playback_100_count"`
	PlayBack25Count  int `json:"playback_25_count"`
	PlayBack50Count  int `json:"playback_50_count"`
	PlayBack75Count  int `json:"playback_75_count"`
	VideoCount       int `json:"video_count"`
}

// MediaPromotedMetrics contains engagement metrics for the media content, tracked in a promoted context, at the time of the request.
type MediaPromotedMetrics struct {
	PlayBack0Count   int `json:"playback_0_count"`
	PlayBack100Count int `json:"playback_100_count"`
	PlayBack25Count  int `json:"playback_25_count"`
	PlayBack50Count  int `json:"playback_50_count"`
	PlayBack75Count  int `json:"playback_75_count"`
	VideoCount       int `json:"video_count"`
}

// MediaPublicMetrics contains public engagement metrics for the media content at the time of the request.
type MediaPublicMetrics struct {
	VideoCount int `json:"video_count"`
}
