package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"groupie/models"
)

var Artists models.Artists

// serve artists.html with artists as struct
func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	if strings.TrimPrefix(r.URL.Path, "/artists/") != "" {
		ErrorHandler(w, "Not found", 404)
		return
	}

	tmp, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		fmt.Println(err)
		ErrorHandler(w, "Internal server error", 500)
		return
	}
	tmp.Execute(w, struct{ Artists models.Artists }{Artists: Artists})
}
