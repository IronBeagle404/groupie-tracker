package utils

import (
	"groupie-tracker/data"
	"log"
	"text/template"
)

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
