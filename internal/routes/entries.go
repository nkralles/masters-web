package routes

import (
	"github.com/nkralles/masters-web/internal/persistence"
	"net/http"
)

func GetEntries(w http.ResponseWriter, r *http.Request) {
	entries, err := persistence.DefaultDriver().GetEntries(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJSON(w, entries, http.StatusOK)
}
