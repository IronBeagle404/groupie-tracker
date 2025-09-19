package utils

import (
	"encoding/json"
	"groupie-tracker/config"
	"groupie-tracker/data"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

// Fetch all data from API
func FetchAllData() {
	fetchData := func(endpoint string, dest interface{}) error {
		res, err := http.Get(config.API_URL + endpoint)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		return json.NewDecoder(res.Body).Decode(dest)
	}
	fetchData("/artists", &data.Artists)
	// fetch locations
	for x, artist := range data.Artists {
		var location data.Location
		fetchData("/locations/"+strconv.Itoa(artist.Id), &location)
		data.Artists[x].Location = location
	}
	// fetch dates
	for x, artist := range data.Artists {
		var dates data.Dates
		fetchData("/dates/"+strconv.Itoa(artist.Id), &dates)
		data.Artists[x].Date = dates
	}
	// fetch relations
	for x, artist := range data.Artists {
		var relation data.Relations
		fetchData("/relation/"+strconv.Itoa(artist.Id), &relation)
		data.Artists[x].Relation = relation
	}
}

func ParseTemplates() {
	var err error
	data.Templates, err = template.ParseGlob("./assets/templates/*.html")
	if err != nil {
		log.Printf("Error parsing templates : %v", err)
	}
}
