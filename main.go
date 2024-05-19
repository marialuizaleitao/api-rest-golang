package main

import (
	"api-rest-golang/database"
	"api-rest-golang/routes"
	"fmt"
)

func main() {
	database.ConnectDatabase()
	fmt.Println("Initializing Rest server with Go")
	routes.HandleRequest()
}
