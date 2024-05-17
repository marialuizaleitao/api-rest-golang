package routes

import (
	"api-rest-golang/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/allPilots", controllers.AllPilots).Methods("Get")
	r.HandleFunc("/allPilots/{id}", controllers.ReturnPilot).Methods("Get")

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		fmt.Println(err)
	}

}
