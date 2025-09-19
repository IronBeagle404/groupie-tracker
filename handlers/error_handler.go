package handlers

import (
	"groupie-tracker/data"
	"groupie-tracker/logging"
	"log"
	"net/http"
)

// Display error with status code, general error type and custom error message
// Write error code to header, log request, and execute error page template
func ErrorHandler(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	tmpl := data.Templates.Lookup("error.html")
	if tmpl == nil {
		log.Print("Error : Template error.html not found")
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(statusCode)
	logging.Logger.Printf("%v \"%v %v %v\" %v", r.RemoteAddr, r.Method, r.URL.Path, r.Proto, statusCode)
	err := tmpl.Execute(w, data.Error{
		Message: message,
		Code:    statusCode,
		Error:   http.StatusText(statusCode),
	})
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
