package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie/handlers"
	"groupie/models"
)

func main() {
	// Get Artists data
	handlers.Artists = models.ParseJson()
	models.AddLocations(handlers.Artists)
	models.AddConcertDates(handlers.Artists)
	models.AddRelations(handlers.Artists)

	// create an empty mux
	mux := http.NewServeMux()

	// handle routes
	mux.HandleFunc("GET /", handlers.NotFound)
	mux.HandleFunc("GET /artists/", handlers.ArtistsHandler)
	mux.HandleFunc("GET /templates/", handlers.ServeStatic)

	// start server
	fmt.Println("Starting server on http://localhost:8081/artists")
	log.Panic(http.ListenAndServe(":8081", mux))
}
