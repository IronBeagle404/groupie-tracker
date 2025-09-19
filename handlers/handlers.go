package handlers

import (
	"groupie-tracker/data"
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
	log.Printf("%v 200 %v", r.Method, r.URL.Path)
}

// Display error with status code, general error type and custom error message
func ErrorHandler(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	tmpl := data.Templates.Lookup("error.html")
	if tmpl == nil {
		log.Print("Error : Template error.html not found")
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}
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
	log.Printf("%v %v %v", r.Method, statusCode, r.URL.Path)
}

// Parse artist ID and display artist page or artist list
func ArtistHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		ErrorHandler(w, r, 405, "Method Not Allowed : Use GET")
		return
	}

	idStr, _ := strings.CutPrefix(r.URL.Path, "/artist/")

	if idStr == "" {
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
	} else {
		idStr, _ = strings.CutSuffix(idStr, "/")
		if strings.Contains(idStr, "/") {
			ErrorHandler(w, r, 404, "Page does not exists")
			return
		}
		fetchId, err := strconv.Atoi(idStr)
		if err != nil || fetchId >= len(data.Artists) || fetchId < 0 {
			ErrorHandler(w, r, 404, "Wrong artist ID")
			return
		}
		tmpl := data.Templates.Lookup("artist.html")
		if tmpl == nil {
			ErrorHandler(w, r, 500, "Template not found")
			return
		}

		var artist data.Artist
		for _, v := range data.Artists {
			if v.Id == fetchId+1 {
				artist.Id = v.Id
				artist.CreationDate = v.CreationDate
				artist.FirstAlbum = v.FirstAlbum
				artist.Image = v.Image
				artist.Members = v.Members
				artist.Name = v.Name
				artist.Relation = v.Relation
			}
		}
		err = tmpl.Execute(w, artist)

		if err != nil {
			ErrorHandler(w, r, 500, "Error rendering template")
			return
		}
	}
	log.Printf("%v 200 %v", r.Method, r.URL.Path)
}
