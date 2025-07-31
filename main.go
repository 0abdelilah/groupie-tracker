package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"groupie/handlers"
	"groupie/models"
)

// move json from file to object
func parseJson(jsonFile string) models.Artists {
	var artists models.Artists

	// Read json
	data, err := os.ReadFile(jsonFile)
	if err != nil {
		fmt.Println("Failed to parse:", err)
		return nil
	}

	// Read and put to struct
	err = json.Unmarshal(data, &artists)
	if err != nil {
		fmt.Println("Failed to parse:", err)
		return nil
	}

	return artists
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func main() {
	// Get Artists data
	handlers.Artists = parseJson("data/artists.json")

	// create an empty mux
	mux := http.NewServeMux()

	// handle routes
	mux.HandleFunc("GET /", NotFound)
	mux.HandleFunc("GET /artists/", handlers.ArtistsHandler)
	mux.HandleFunc("GET /locations/{id}", handlers.LocationsHandler)
	mux.HandleFunc("GET /relation/{id}", handlers.RelationHandler)
	mux.HandleFunc("GET /dates/{id}", handlers.DatesHandler)

	// serve static files (css, js)
	mux.HandleFunc("GET /templates/", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))).ServeHTTP(w, r)
	})

	// start server
	fmt.Println("Starting server on http://localhost:8080/artists")
	log.Panic(http.ListenAndServe(":8080", mux))
}
