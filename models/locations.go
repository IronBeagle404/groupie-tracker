package models

type Location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Locations struct {
	Index []Location `json:"index"`
}
