package persistence

import "time"

type Score struct {
	PlayerId    int       `json:"playerId,omitempty"`
	Rank        int       `json:"rank,omitempty"`
	FirstName   string    `json:"first_name,omitempty"`
	LastName    string    `json:"last_name,omitempty"`
	CountryCode string    `json:"countryCode"`
	Rounds      []Round   `json:"rounds"`
	ToPar       int       `json:"total"`
	LastUpdated time.Time `json:"lastUpdated"`
	Standing    int       `json:"standing"`
}

type Round struct {
	Round       int       `json:"round"`
	ToPar       int       `json:"toPar"`
	LastUpdated time.Time `json:"lastUpdated"`
}
