package handlers

import (
	"io"
	"net/http"
)

// fetch locations & return it
func LocationsHandler(w http.ResponseWriter, r *http.Request) {
	id := GetID(r, "locations")
	if id == "" {
		http.NotFound(w, r)
		return
	}

	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + id)
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
