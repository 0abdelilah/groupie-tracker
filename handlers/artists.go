package handlers

import (
	"groupie/models"
	"html/template"
	"net/http"
)

var Artists models.Artists

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("templates/artists.html")
	tmp.Execute(w, struct{ Artists models.Artists }{Artists: Artists})
}
