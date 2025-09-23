package main

import (
	"groupie-tracker/config"
	"groupie-tracker/data"
	"groupie-tracker/handlers"
	"groupie-tracker/logging"
	"groupie-tracker/utils"
	"net/http"
	"strings"
)

// Initialize logging, attempt to fetch all data from API and starting server if successful
func init() {
	var err error
	logging.Init()
	logging.Logger.Print("Starting server...")
	data.CombinedData, err = utils.FetchAllData()
	if err != nil {
		logging.Logger.Printf("Failed to fetch data : %v", err)
	}
	if len(data.CombinedData.Artists) == 0 {
		logging.Logger.Print("Data was not successfully fetched during server init. Retrying...")
		data.CombinedData, err = utils.FetchAllData()
		if err != nil {
			logging.Logger.Fatalf("Failed to fetch data : %v", err)
		}
		logging.Logger.Print("Data fetched successfully after retry")
	}
}

// Handle assets, parse all templates, handle routes, and start server
func main() {
	mux := http.NewServeMux()

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
	mux.HandleFunc("/search", handlers.SearchHandler)
	mux.HandleFunc("/filter", handlers.FilterHandler)

	logging.Logger.Print("Server successfully started at http://localhost:" + config.PORT)
	logging.Logger.Fatal(http.ListenAndServe(":"+config.PORT, mux))
}
