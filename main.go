package main

import (
	"api-rest-golang/models"
	"api-rest-golang/routes"
	"fmt"
)

func main() {
	models.Pilots = []models.Pilot{
		{Id: 1, Name: "Ayrton Senna", History: "Ayrton Senna was a Brazilian racing driver who won three Formula One world championships for McLaren in 1988, 1990, and 1991. He is considered one of the greatest Formula 1 drivers of all time due to his exceptional speed, skill, and intense competitive spirit. Senna's tragic death at the 1994 San Marino Grand Prix shocked the world and left an indelible mark on the sport."},
		{Id: 2, Name: "Emerson Fittipaldi",
			History: "Emerson Fittipaldi is a Brazilian racing driver who became the youngest Formula One world champion at the time when he won the title in 1972 at the age of 25. He won his second world championship in 1974. Fittipaldi is also known for his success in other racing series, including CART and IndyCar. He remains one of Brazil's most successful racing drivers."},
	}

	fmt.Println("Initializing Rest server with Go")
	routes.HandleRequest()
}
