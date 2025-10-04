package handlers

import (
	"net/http"
	"os"
)

func ServeStatic(w http.ResponseWriter, r *http.Request) {
	bytes, err := os.ReadFile("." + r.URL.Path)
	if err != nil {
		ErrorHandler(w, "This page does not exist.", 404)
	}

	w.Write(bytes)
}
