package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func RunServer() {
	log.Print("Starting server at http://localhost:8080")
	http.Handle("/assets/static/", http.StripPrefix("/assets/static/", http.FileServer(http.Dir("assets/static"))))
	http.HandleFunc("/", home)

	http.HandleFunc("GET /artist/", artistHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Print(r.Method, " ", http.StatusNotFound, " ", r.URL.Path)
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles("./assets/templates/index.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Error : Invalid Template", http.StatusNotFound)
		return
	}
	err = tmpl.Execute(w, "")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("%v 200 %v", r.Method, r.URL.Path)
}

func getArtists() []Artist {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err.Error())
		return nil
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err.Error())
		return nil
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Print(err.Error())
		return nil
	}

	var artists []Artist
	err = json.Unmarshal([]byte(body), &artists)
	if err != nil {
		log.Print(err.Error())
		return nil
	}

	return artists
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
	idStr, _ := strings.CutPrefix(r.URL.Path, "/artist/")
	Artists := getArtists()

	if idStr == "" {
		tmpl, err := template.ParseFiles("./assets/templates/artistlist.html")
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Error : Invalid Template", http.StatusNotFound)
			return
		}
		err = tmpl.Execute(w, Artists)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		idStr, _ = strings.CutSuffix(idStr, "/")
		if strings.Contains(idStr, "/") {
			http.Error(w, "Error : 404", http.StatusNotFound)
			log.Print(r.Method, " ", http.StatusNotFound, " ", r.URL.Path)
			return
		}
		fetchId, err := strconv.Atoi(idStr)
		if err != nil || fetchId >= len(Artists) || fetchId < 0 {
			http.Error(w, "Error : Invalid Artist ID", http.StatusNotFound)
			log.Print(r.Method, " ", http.StatusNotFound, " ", r.URL.Path)
			return
		}
		tmpl, err := template.ParseFiles("./assets/templates/artist.html")
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Error : Invalid Template", http.StatusNotFound)
			return
		}
		err = tmpl.Execute(w, Artists[fetchId])
		if err != nil {
			log.Print(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	log.Printf("%v 200 %v", r.Method, r.URL.Path)
}
