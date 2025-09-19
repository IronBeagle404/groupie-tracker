package main

import (
	"groupie-tracker/config"
	"groupie-tracker/data"
	"groupie-tracker/handlers"
	"groupie-tracker/logging"
	"groupie-tracker/utils"
	"log"
	"net/http"
	"strings"
)

// Attempt to fetch all data from API and starting server if successful
func init() {
	var err error
	data.CombinedData, err = utils.FetchAllData()
	if err != nil {
		log.Fatalf("Failed to fetch data : %v", err)
	}
	if len(data.CombinedData.Artists) == 0 {
		log.Fatal("Data was not successfully fetched during server init, closing server...")
	}
}

// Initialize logging, handle assets, parse all templates, handle routes, and start server
func main() {
	mux := http.NewServeMux()
	logging.Init()

	mux.Handle("/assets/static/", http.StripPrefix("/assets/static/", http.FileServer(http.Dir("assets/static"))))
	utils.ParseTemplates()

	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/artist/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/artist/")
		if id == "" {
			handlers.ArtistListHandler(w, r)
			return
		}
		handlers.ArtistHandler(w, r)
	})
	mux.HandleFunc("/healthz", handlers.HealthCheck)

	logging.Logger.Print("Starting server at http://localhost:" + config.PORT)
	logging.Logger.Fatal(http.ListenAndServe(":"+config.PORT, mux))
}
