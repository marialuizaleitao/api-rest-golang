package models

type Pilot struct {
	Name    string `json:"name"`
	History string `json:"history"`
}

var Pilots []Pilot
