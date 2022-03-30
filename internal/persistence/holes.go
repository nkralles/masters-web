package persistence

type Hole struct {
	Hole  int    `json:"hole,omitempty"`
	Name  string `json:"name,omitempty"`
	Par   int    `json:"par,omitempty"`
	Yards int    `json:"yards,omitempty"`
}
