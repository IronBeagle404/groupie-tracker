package handlers

import (
	"groupie-tracker/data"
	"groupie-tracker/logging"
	"groupie-tracker/models"
	"groupie-tracker/utils"
	"net/http"
	"slices"
	"strconv"
)

func FilterHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		ErrorHandler(w, r, 405, "Method Not Allowed : Use GET")
		return
	}

	var (
		CreationDateStart int
		CreationDateEnd   int
		FirstAlbumStart   int
		FirstAlbumEnd     int
		Location          string
		MembersStr        []string
		Members           []int
		Ids               []int
		err               error
	)

	// if all filters are empty, display all artists
	r.ParseForm()
	filledFilters := 0
	for _, filter := range r.Form {
		if filter[0] != "" {
			filledFilters++
		}
	}
	if filledFilters == 0 {
		for x := 0; x <= len(data.Artists); x++ {
			a, _ := utils.FetchArtist(x)
			Ids = append(Ids, a.Id)
		}
	}

	// parse creationDate & firstAlbum filter
	CreationDateStart, err = strconv.Atoi(r.FormValue("CreationDateStart"))
	if err != nil {
		CreationDateStart = 1945
	}
	CreationDateEnd, err = strconv.Atoi(r.FormValue("CreationDateEnd"))
	if err != nil {
		CreationDateEnd = 2025
	}
	if CreationDateStart > CreationDateEnd {
		CreationDateStart, CreationDateEnd = CreationDateEnd, CreationDateStart
	}
	FirstAlbumStart, err = strconv.Atoi(r.FormValue("FirstAlbumStart"))
	if err != nil {
		FirstAlbumStart = 1945
	}
	FirstAlbumEnd, err = strconv.Atoi(r.FormValue("FirstAlbumEnd"))
	if err != nil {
		FirstAlbumEnd = 2025
	}
	if FirstAlbumStart > FirstAlbumEnd {
		FirstAlbumStart, FirstAlbumEnd = FirstAlbumEnd, FirstAlbumStart
	}

	// parse Location filter
	// no presence/error check should be necessary
	Location = r.FormValue("location")

	// parse members if present
	MembersStr = r.URL.Query()["members"]
	for _, str := range MembersStr {
		nbr, err := strconv.Atoi(str)
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError, "Error during filter parsing")
			return
		}
		Members = append(Members, nbr)
	}

	// range over stored data and add ids checking filters conditions
	for _, a := range data.Artists {
		ArtistFirstAlbum, _ := strconv.Atoi(a.FirstAlbum[6:])

		if !(a.CreationDate >= CreationDateStart && a.CreationDate <= CreationDateEnd) {
			continue
		}

		if !(ArtistFirstAlbum >= FirstAlbumStart && ArtistFirstAlbum <= FirstAlbumEnd) {
			continue
		}

		if Location != "" {
			locFound := false
			for _, r := range data.Relations.Index {
				if r.Id == a.Id {
					for loc := range r.DatesLocations {
						if loc == Location {
							locFound = true
							break
						}
					}
				}
			}
			if !locFound {
				continue
			}
		}

		if Members != nil {
			validLen := false
			for _, n := range Members {
				if len(a.Members) == n {
					validLen = true
				}
			}
			if !validLen {
				continue
			}
		}

		if slices.Contains(Ids, a.Id) {
			continue
		}

		Ids = append(Ids, a.Id)
	}

	var FilteredData models.Output
	FilteredData.For_Search = data.CombinedData

	// gather data for selected ids
	for _, id := range Ids {
		new, err := utils.FetchArtist(id)
		if err == nil {
			FilteredData.To_Display.Artists = append(FilteredData.To_Display.Artists, new)
		}
	}

	// fetch and fill template with filtered data
	tmpl := data.Templates.Lookup("artistlist.html")
	if tmpl == nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Template not found")
		return
	}
	err = tmpl.Execute(w, FilteredData)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Error rendering template")
		return
	}

	// log success
	logging.Logger.Printf("%v \"%v %v %v\" %v", r.RemoteAddr, r.Method, r.URL.String(), r.Proto, http.StatusOK)
}
