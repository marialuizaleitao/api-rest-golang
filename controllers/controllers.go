package controllers

import (
	"api-rest-golang/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func AllPilots(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Pilots)
}

func ReturnPilot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, pilot := range models.Pilots {
		if strconv.Itoa(pilot.Id) == id {
			json.NewEncoder(w).Encode(pilot)
		}
	}
}
