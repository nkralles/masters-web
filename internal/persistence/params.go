package persistence

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

type PagingParams struct {
	Offset    int64  `json:"offset,omitempty"`
	Limit     int64  `json:"limit,omitempty"`
	SortBy    string `json:"sortBy,omitempty"`
	Ascending bool   `json:"ascending"`
}

type TextParams struct {
	Query string `json:"query,omitempty"`
}

func ParseTextParams(r *http.Request) (params *TextParams) {
	r.ParseForm()
	return &TextParams{
		Query: r.Form.Get("q"),
	}
}

// ParsePagingParams parses the standard paging params
func ParsePagingParams(r *http.Request, sort string) (params *PagingParams, err error) {
	r.ParseForm()
	sortBy := sort
	offset := 0
	if r.Form.Get("offset") != "" {
		var v int
		v, err = strconv.Atoi(r.Form.Get("offset"))
		if err == nil && v >= 0 {
			offset = v
		} else if err != nil {
			return
		} else {
			err = errors.New("invalid offset")
			return
		}
	}
	limit := 0
	if r.Form.Get("limit") != "" {
		var v int
		v, err = strconv.Atoi(r.Form.Get("limit"))
		if err == nil && v >= -1 {
			limit = v
		} else if err != nil {
			return
		} else {
			err = errors.New("invalid limit")
			return
		}
	}
	if v := r.Form.Get("sortBy"); v != "" {
		sortBy = v
	}
	ascending := true
	if v := r.Form.Get("ascending"); strings.EqualFold(v, "false") {
		ascending = false
	}
	params = &PagingParams{
		Offset:    int64(offset),
		Limit:     int64(limit),
		SortBy:    sortBy,
		Ascending: ascending,
	}
	return
}

type CommonParams struct {
	*PagingParams
	*TextParams
}
