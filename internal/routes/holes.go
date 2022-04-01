package routes

import (
	"github.com/gorilla/mux"
	"github.com/nkralles/masters-web/internal/persistence"
	"net/http"
	"strconv"
)

func GetHole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hole, err := strconv.Atoi(vars["hole"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	h, err := persistence.DefaultDriver().GetHole(r.Context(), hole)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	sendJSON(w, h, http.StatusOK)
}
func GetHoles(w http.ResponseWriter, r *http.Request) {
	holes, err := persistence.DefaultDriver().GetHoles(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	sendJSON(w, holes, http.StatusOK)
}
