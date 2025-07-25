package main

import (
	"encoding/json"
	"fmt"
	"groupie/handlers"
	"groupie/models"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

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

func main() {

	// Get Artists data
	handlers.Artists = parseJson("data/artists.json")

	// create an empty mux
	mux := http.NewServeMux()

	// handle routes
	mux.HandleFunc("GET /", handlers.ArtistsHandler)
	mux.HandleFunc("GET /locations/{id}", handlers.LocationsHandler)
	mux.HandleFunc("GET /relation/{id}", relationHandler)

	// serve static files (css, js)
	mux.HandleFunc("GET /templates/", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))).ServeHTTP(w, r)
	})

	// start server
	fmt.Println("Starting server on http://localhost:8080/artists")
	log.Panic(http.ListenAndServe(":8080", mux))
}

func relationHandler(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/relation/")
	if id == "" {
		http.NotFound(w, r)
		return
	}

	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + id)
	if err != nil || resp.StatusCode != 200 {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)

}
