package routes

import (
	"github.com/nkralles/masters-web/internal/persistence"
	"net/http"
)

func GetHoles(w http.ResponseWriter, r *http.Request) {
	holes, err := persistence.DefaultDriver().GetHoles(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	sendJSON(w, holes, http.StatusOK)
}
