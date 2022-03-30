package routes

import (
	"encoding/json"
	"net/http"
)

const (
	// ContentTypeHeader is the name of the HTTP Content-Type header
	ContentTypeHeader = "Content-Type"
)

// Content-Type header values
const (
	// JSONContentType is the MIME Type for JSON files
	JSONContentType = "application/json"
)

func sendJSON(w http.ResponseWriter, doc interface{}, status int) {
	b, err := json.Marshal(doc)
	if err != nil {
		http.Error(w, "Unexpected error while encoding result", http.StatusInternalServerError)
		return
	}
	w.Header().Set(ContentTypeHeader, JSONContentType)
	w.WriteHeader(status)
	w.Write(b)
}
