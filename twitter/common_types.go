package twitter

// Includes contains returned referenced objects if you include an expansion parameter
type Includes struct {
	Tweets []Tweet `json:"tweets,omitempty"`
	Users  []User  `json:"users,omitempty"`
	Places []Place `json:"places,omitempty"`
	Media  []Media `json:"media,omitempty"`
	Polls  []Poll  `json:"polls,omitempty"`
}