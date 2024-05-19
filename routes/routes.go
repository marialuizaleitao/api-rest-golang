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
	r.HandleFunc("/pilots", controllers.AllPilots).Methods("Get")
	r.HandleFunc("/pilots/{id}", controllers.ReturnPilot).Methods("Get")
	r.HandleFunc("/pilots", controllers.CreatePilot).Methods("Post")
	r.HandleFunc("/pilots/{id}", controllers.DeletePilot).Methods("Delete")
	r.HandleFunc("/pilots/{id}", controllers.UpdatePilot).Methods("Put")

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		fmt.Println(err)
	}

}
