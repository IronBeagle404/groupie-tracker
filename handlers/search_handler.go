package handlers

import (
	"groupie-tracker/data"
	"groupie-tracker/logging"
	"groupie-tracker/models"
	"groupie-tracker/utils"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		ErrorHandler(w, r, 405, "Method Not Allowed : Use GET")
		return
	}

	var Ids []int

	key := r.FormValue("Search")
	// harcode key for testing
	//key := ""

	key = strings.ToLower(key)

	for _, a := range data.Artists {
		if strings.Contains(strings.ToLower(a.Name), key) || a.FirstAlbum == key || strconv.Itoa(a.CreationDate) == key {
			if !slices.Contains(Ids, a.Id) {
				Ids = append(Ids, a.Id)
			}
		}

		for _, m := range a.Members {
			if strings.Contains(strings.ToLower(m), key) {
				if !slices.Contains(Ids, a.Id) {
					Ids = append(Ids, a.Id)
				}
			}
		}
	}

	for _, i := range data.Locations.Index {
		for _, l := range i.Locations {
			if strings.Contains(strings.ToLower(l), key) {
				if !slices.Contains(Ids, i.Id) {
					Ids = append(Ids, i.Id)
				}
			}
		}
	}

	var SearchData models.Output

	// Add artist IDs if fetched successfuly
	for _, id := range Ids {
		new, err := utils.FetchArtist(id)
		if err == nil {
			SearchData.To_Display.Artists = append(SearchData.To_Display.Artists, new)
		}
	}

	SearchData.For_Search = data.CombinedData

	tmpl := data.Templates.Lookup("artistlist.html")
	if tmpl == nil {
		ErrorHandler(w, r, 500, "Template not found")
		return
	}
	err := tmpl.Execute(w, SearchData)
	if err != nil {
		ErrorHandler(w, r, 500, "Error rendering template")
		return
	}

	logging.Logger.Printf("%v \"%v %v %v\" %v", r.RemoteAddr, r.Method, r.URL.String(), r.Proto, http.StatusOK)
}
