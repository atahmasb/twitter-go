package twitter

// Place contains information about a place tagged in a Tweet.
type Place struct {
	// A longer-form detailed place name.
	FullName string `json:"full_name,required"`

	// The unique identifier of the expanded place,
	// if this is a point of interest tagged in the Tweet.
	ID string `json:"id,required"`

	// Returns the identifiers of known places that contain
	// the referenced place.
	ContainedWithin []string `json:"contained_within"`

	// The full-length name of the country this place belongs to.
	Country string `json:"country"`

	// The ISO Alpha-2 country code this place belongs to.
	CountryCode string `json:"country_code"`

	// Contains place details in GeoJSON format.
	Geo PlaceGeoInfo `json:"geo"`

	// The short name of this place.
	Name string `json:"name"`

	// Specified the particular type of information represented by
	// this place information, such as a city name, or a point of interest.
	PlaceType string `json:"place_type"`
}

// PlaceGeoInfo contains place details in GeoJSON format.
type PlaceGeoInfo struct {
	Type       string    `json:"type"`
	Bbox       []float64 `json:"bbox"`
	Properties string    `json:"properties"`
}
