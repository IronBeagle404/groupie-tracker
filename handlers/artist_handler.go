package handlers

import (
	"fmt"
	"groupie-tracker/data"
	"groupie-tracker/logging"
	"groupie-tracker/models"
	"groupie-tracker/utils"
	"net/http"
	"strconv"
	"strings"
)

// Display list of all artists in the API
func ArtistListHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		ErrorHandler(w, r, 405, "Method Not Allowed : Use GET")
		return
	}

	tmpl := data.Templates.Lookup("artistlist.html")
	if tmpl == nil {
		ErrorHandler(w, r, 500, "Template not found")
		return
	}

	allData := models.Output{
		To_Display: data.CombinedData,
		For_Search: data.CombinedData,
	}

	err := tmpl.Execute(w, allData)
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
	if err != nil {
		msg := fmt.Sprintf("\"%v\" is not a valid ID", idStr)
		ErrorHandler(w, r, http.StatusBadRequest, msg)
		return
	}
	tmpl := data.Templates.Lookup("artist.html")
	if tmpl == nil {
		ErrorHandler(w, r, 500, "Template not found")
		return
	}

	artist, err := utils.FetchArtist(fetchId)
	if err != nil {
		ErrorHandler(w, r, http.StatusNotFound, err.Error())
		return
	}
	err = tmpl.Execute(w, artist)
	if err != nil {
		ErrorHandler(w, r, 500, "Error rendering template")
		return
	}
	logging.Logger.Printf("%v \"%v %v %v\" %v", r.RemoteAddr, r.Method, r.URL.Path, r.Proto, http.StatusOK)
}
