package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`

	LocationsPath string   `json:"locations"`
	Locations     []string `json:"-"`

	ConcertDatesPath string   `json:"concertDates"`
	ConcertDates     []string `json:"-"`

	RelationsPath string `json:"relations"`
	Relations     string `json:"-"`
}

type Artists []Artist

/*
loop over each json
fetch
add locations to the struct
*/
func AddLocations(artists Artists) {
	for i, artist := range artists {
		resp, err := http.Get(artist.LocationsPath)
		if err != nil {
			panic(err)
		}

		var location struct {
			Locations []string `json:"locations"`
		}

		jsonData, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(jsonData, &location)
		if err != nil {
			panic(err)
		}

		artists[i].Locations = location.Locations
	}
}

func AddConcertDates(artists Artists) {
	for i, artist := range artists {
		resp, err := http.Get(artist.ConcertDatesPath)
		if err != nil {
			panic(err)
		}

		var ConcertDates struct {
			Dates []string `json:"dates"`
		}

		jsonData, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(jsonData, &ConcertDates)
		if err != nil {
			panic(err)
		}

		artists[i].ConcertDates = ConcertDates.Dates
	}
}

func AddRelations(artists Artists) {
	for i, artist := range artists {
		resp, err := http.Get(artist.RelationsPath)
		if err != nil {
			panic(err)
		}

		jsonData, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			panic(err)
		}

		var relations struct {
			DatesLocations map[string][]string `json:"datesLocations"`
		}

		err = json.Unmarshal(jsonData, &relations)
		if err != nil {
			panic(err)
		}

		var sb strings.Builder
		for city, dates := range relations.DatesLocations {
			d := strings.Join(dates, ", ")
			fmt.Fprintf(&sb, "%s => %s\n", city, d)
		}

		artists[i].Relations = sb.String()
	}
}
