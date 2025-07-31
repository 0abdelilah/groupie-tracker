package handlers

import (
	"html/template"
	"net/http"
	"strings"

	"groupie/models"
)

var Artists models.Artists

// serve artists.html with artists as struct
func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	if strings.TrimPrefix(r.URL.Path, "/artists/") != "" {
		http.NotFound(w, r)
		return
	}

	tmp, _ := template.ParseFiles("templates/artists.html")
	tmp.Execute(w, struct{ Artists models.Artists }{Artists: Artists})
}
