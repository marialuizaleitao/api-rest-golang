package controllers

import (
	"api-rest-golang/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func AllPilots(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Pilots)
}
