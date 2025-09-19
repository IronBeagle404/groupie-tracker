package models

type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Dates struct {
	Index []Date `json:"index"`
}
