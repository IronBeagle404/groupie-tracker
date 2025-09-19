package main

import (
	"groupie-tracker/config"
	"groupie-tracker/handlers"
	"groupie-tracker/utils"
	"log"
	"net/http"
)

// Fetch all data from API before starting server
func init() {
	utils.FetchAllData()
}

// Start and run server
func main() {
	log.Print("Starting server at http://localhost:" + config.PORT)
	http.Handle("/assets/static/", http.StripPrefix("/assets/static/", http.FileServer(http.Dir("assets/static"))))

	utils.ParseTemplates()

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/artist/", handlers.ArtistHandler)

	log.Fatal(http.ListenAndServe(":"+config.PORT, nil))
}
