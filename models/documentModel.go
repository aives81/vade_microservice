package models

type Document struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var Documents = make(map[string]Document)
