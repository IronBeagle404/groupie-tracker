package handlers

import (
	"groupie-tracker/data"
	"groupie-tracker/logging"
	"groupie-tracker/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
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

// Display list of all artists in the API
func ArtistListHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := data.Templates.Lookup("artistlist.html")
	if tmpl == nil {
		ErrorHandler(w, r, 500, "Template not found")
		return
	}
	err := tmpl.Execute(w, data.Artists)
	if err != nil {
		ErrorHandler(w, r, 500, "Error rendering template")
		return
	}
	logging.Logger.Printf("%v \"%v %v %v\" %v", r.RemoteAddr, r.Method, r.URL.Path, r.Proto, http.StatusOK)
}

// Parse artist ID, fetch corresponding data and display the artist page
func ArtistHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		ErrorHandler(w, r, 405, "Method Not Allowed : Use GET")
		return
	}

	idStr, _ := strings.CutPrefix(r.URL.Path, "/artist/")

	idStr, _ = strings.CutSuffix(idStr, "/")
	if strings.Contains(idStr, "/") {
		ErrorHandler(w, r, 404, "Page does not exists")
		return
	}
	fetchId, err := strconv.Atoi(idStr)
	if err != nil || fetchId > 52 || fetchId < 1 {
		ErrorHandler(w, r, 404, "Wrong artist ID")
		return
	}
	tmpl := data.Templates.Lookup("artist.html")
	if tmpl == nil {
		ErrorHandler(w, r, 500, "Template not found")
		return
	}

	artist := utils.FetchArtist(fetchId)
	err = tmpl.Execute(w, artist)

	if err != nil {
		ErrorHandler(w, r, 500, "Error rendering template")
		return
	}
	logging.Logger.Printf("%v \"%v %v %v\" %v", r.RemoteAddr, r.Method, r.URL.Path, r.Proto, http.StatusOK)
}

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
