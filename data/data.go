package data

import (
	"groupie-tracker/models"
	"text/template"
)

var (
	Artists      []models.Artist
	Dates        models.Dates
	Locations    models.Locations
	Relations    models.Relations
	CombinedData models.CombinedData
)

var Templates *template.Template

type Error struct {
	Message string
	Code    int
	Error   string
}
