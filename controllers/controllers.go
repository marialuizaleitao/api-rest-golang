package controllers

import (
	"api-rest-golang/database"
	"api-rest-golang/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func AllPilots(w http.ResponseWriter, r *http.Request) {
	var p []models.Pilot
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func ReturnPilot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Pilot
	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}

func CreatePilot(w http.ResponseWriter, r *http.Request) {
	var newPilot models.Pilot
	json.NewDecoder(r.Body).Decode(&newPilot)
	database.DB.Create(&newPilot)
	json.NewEncoder(w).Encode(newPilot)
}
