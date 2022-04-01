package routes

import (
	"github.com/gorilla/mux"
	"github.com/nkralles/masters-web/internal/persistence"
	"net/http"
	"strconv"
)

func GetGolfer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.Atoi(vars["player_id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	golfer, err := persistence.DefaultDriver().GetGolferById(r.Context(), pid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sendJSON(w, golfer, http.StatusOK)
	return
}
func GetGolfers(w http.ResponseWriter, r *http.Request) {
	text := persistence.ParseTextParams(r)
	paging, err := persistence.ParsePagingParams(r, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	golfers, err := persistence.DefaultDriver().GetGolfers(r.Context(), &persistence.CommonParams{
		PagingParams: paging,
		TextParams:   text,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sendJSON(w, golfers, http.StatusOK)
	return
}
