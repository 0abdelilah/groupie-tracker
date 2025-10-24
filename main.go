package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie/handlers"
	"groupie/models"
)

// todo:
// make popups look better
// add shortcuts

func main() {
	// Get Artists data
	handlers.Artists = models.ParseJson()
	models.AddLocations(handlers.Artists)
	models.AddDates(handlers.Artists)
	models.AddRelations(handlers.Artists)

	// create an empty mux
	mux := http.NewServeMux()

	// handle routes
	mux.HandleFunc("GET /", handlers.ArtistsHandler)
	mux.HandleFunc("GET /templates/", handlers.ServeStatic)

	// start server
	fmt.Println("Starting server on http://localhost:8081/")
	log.Panic(http.ListenAndServe(":8081", mux))
}
