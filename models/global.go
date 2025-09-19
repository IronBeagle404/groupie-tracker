package models

type CombinedData struct {
	Artists   []Artist
	Locations []Location
	Dates     []Date
	Relations []Relation
}

type Output struct {
	To_Display CombinedData
	For_Search CombinedData
}
