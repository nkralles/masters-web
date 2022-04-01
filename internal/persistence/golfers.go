package persistence

type Golfer struct {
	PlayerId    int    `json:"playerId,omitempty"`
	Rank        int    `json:"rank,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	CountryCode string `json:"cc,omitempty"`
}

func (g *Golfer) Top12() bool {
	return g.Rank <= 12
}

type GolferResponse struct {
	Golfers *[]Golfer `json:"golfers,omitempty"`
	PagingResponse
}
