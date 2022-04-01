package persistence

type PagingResponse struct {
	Total int64 `json:"total,omitempty"`
	Count int64 `json:"count,omitempty"`
}
