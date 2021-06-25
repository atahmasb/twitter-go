package twitter

import (
	"time"
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

// Poll can be included in a Tweet is not a primary object on any endpoint,
// but can be found and expanded in the Tweet object.
type Poll struct {
	// Unique identifier of the expanded poll.
	ID string `json:"id"`

	// Contains objects describing each choice in the referenced poll.
	Options []PollOption `json:"options"`

	// Specifies the total duration of this poll in minutes.
	Duration int `json:"duration_minute"`

	// Specifies the end date and time for this poll.
	EndDateTime time.Time

	// Indicates if this poll is still active and can receive votes, or if the voting is now closed.
	VotingStatus string `json:"voting_status"`
}

// PollOption contains details of an option in a Poll.
type PollOption struct {
	Position int    `json:"position"`
	Label    string `json:"label"`
	Votes    int    `json:"votes"`
}
