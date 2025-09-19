package handlers

import (
	"groupie-tracker/logging"
	"net/http"
)

// Check server status
// Returns 204 if everything's okay
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, r, 405, "Method Not Allowed : Use GET")
		return
	}
	w.WriteHeader(http.StatusNoContent)
	logging.Logger.Printf("%v \"%v %v %v\" %v", r.RemoteAddr, r.Method, r.URL.Path, r.Proto, http.StatusNoContent)
}
