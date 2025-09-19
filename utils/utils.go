package utils

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/config"
	"groupie-tracker/data"
	"groupie-tracker/models"
	"log"
	"net/http"
	"sync"
	"text/template"
)

// Fetch all data from API
func FetchAllData() (models.CombinedData, error) {
	var (
		wg  sync.WaitGroup
		err error
	)

	fetchData := func(endpoint string, dest interface{}) {
		defer wg.Done()
		res, _ := http.Get(config.API_URL + endpoint)
		if res.StatusCode != http.StatusOK {
			err = fmt.Errorf("API returned status code %d", res.StatusCode)
			return
		}
		defer res.Body.Close()

		json.NewDecoder(res.Body).Decode(dest)
	}

	wg.Add(4)
	go fetchData("/artists", &data.Artists)
	go fetchData("/locations", &data.Locations)
	go fetchData("/dates", &data.Dates)
	go fetchData("/relation", &data.Relations)
	wg.Wait()

	if err != nil {
		return models.CombinedData{}, err
	}

	return models.CombinedData{
		Artists:   data.Artists,
		Locations: data.Locations.Index,
		Dates:     data.Dates.Index,
		Relations: data.Relations.Index,
	}, nil
}

func ParseTemplates() {
	var err error
	funcMap := template.FuncMap{
		"add": func(a, b int) int { return a + b },
	}
	data.Templates, err = template.New("").Funcs(funcMap).ParseGlob("./assets/templates/*.html")
	if err != nil {
		log.Printf("Error parsing templates : %v", err)
	}
}

func FetchArtist(id int) models.Artist {
	var artist models.Artist

	for _, v := range data.Artists {
		if v.Id == id {
			artist.Id = v.Id
			artist.CreationDate = v.CreationDate
			artist.FirstAlbum = v.FirstAlbum
			artist.Image = v.Image
			artist.Members = v.Members
			artist.Name = v.Name
		}
	}

	var location models.Location
	for _, loc := range data.Locations.Index {
		if loc.Id == id {
			location.Locations = loc.Locations
		}
	}

	var date models.Date
	for _, d := range data.Dates.Index {
		if d.Id == id {
			date.Dates = d.Dates
		}
	}

	var relation models.Relation
	for _, rel := range data.Relations.Index {
		if rel.Id == id {
			relation.DatesLocations = rel.DatesLocations
		}
	}

	artist.Location = location
	artist.Date = date
	artist.Relation = relation
	return artist
}
