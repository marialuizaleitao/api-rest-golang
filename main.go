package main

import (
	"api-rest-golang/routes"
	"fmt"
)

func main() {
	fmt.Println("Initializing Rest server with Go")
	routes.HandleRequest()
}
