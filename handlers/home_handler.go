package handlers

import (
	"groupie-tracker/data"
	"groupie-tracker/logging"
	"net/http"
)

// Display homepage
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, r, 404, "Page does not exist")
		return
	}

	if r.Method != http.MethodGet {
		ErrorHandler(w, r, 405, "Method Not Allowed : Use GET")
		return
	}

	tmpl := data.Templates.Lookup("index.html")
	if tmpl == nil {
		ErrorHandler(w, r, 500, "Template not found")
		return
	}
	err := tmpl.Execute(w, nil)
	if err != nil {
		ErrorHandler(w, r, 500, "Error rendering template")
		return
	}
	logging.Logger.Printf("%v \"%v %v %v\" %v", r.RemoteAddr, r.Method, r.URL.Path, r.Proto, http.StatusOK)
}
