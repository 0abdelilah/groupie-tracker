package handlers

import (
	"io"
	"net/http"
)

// fetch json ny id & return it
func RelationHandler(w http.ResponseWriter, r *http.Request) {
	id := GetID(r, "relation")
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
