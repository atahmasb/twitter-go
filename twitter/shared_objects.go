package twitter

// Cashtag is a cashtag found in the text of Tweet.
type Cashtag struct {
	Start int    `json:"start"`
	End   int    `json:"end"`
	Tag   string `json:"tag"`
}

// Hashtag is a hashtag found in the text of Tweet
type Hashtag struct {
	Start int    `json:"start"`
	End   int    `json:"end"`
	Tag   string `json:"tag"`
}

// Mention is a mention found in the text of Tweet
type Mention struct {
	Start int    `json:"start"`
	End   int    `json:"end"`
	Tag   string `json:"tag"`
}

// URL is a URL found in the text of Tweet
type URL struct {
	Start       int    `json:"start"`
	End         int    `json:"end"`
	URL         string `json:"tag"`
	Description string `json:"description"`
}

// Withheld contains withholding details.
type Withheld struct {
	CopyRight   bool     `json:"copyright"`
	ContryCodes []string `json:"country_codes"`
}
