package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"groupie/models"
)

type TemplateData struct {
	Artists models.Artists
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	keyword := r.FormValue("q")

	matchingArtists := Search(w, Artists, keyword)

	tmpl, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	data := TemplateData{
		Artists: matchingArtists,
	}

	tmpl.Execute(w, data)
}

func Search(w http.ResponseWriter, artists models.Artists, keyword string) models.Artists {
	var result models.Artists

	keyword = strings.ToLower(keyword)

	for _, artist := range artists {
		// Search Names
		if strings.Contains(strings.ToLower(artist.Name), keyword) {
			fmt.Printf("Found %s in Names\n", keyword)
			result = append(result, artist)
		}

		// Search Members
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), keyword) {
				fmt.Printf("Found %s in Members\n", keyword)
				result = append(result, artist)
			}
		}
	}
	return result
}
