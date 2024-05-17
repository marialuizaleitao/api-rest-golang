package models

type Pilot struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	History string `json:"history"`
}

var Pilots []Pilot
