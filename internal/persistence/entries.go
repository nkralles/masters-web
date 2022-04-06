package persistence

type Entry struct {
	Name         string    `json:"name,omitempty"`
	WinningScore int       `json:"winning_score"`
	Golfers      *[]Golfer `json:"golfers"`
	Total        int       `json:"total"`
	Rank         int       `json:"rank"`
}
