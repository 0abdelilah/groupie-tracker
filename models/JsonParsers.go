package models

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type Artist struct {
	Id            int      `json:"id"`
	Image         string   `json:"image"`
	Name          string   `json:"name"`
	Members       []string `json:"members"`
	JoinedMembers string   `json:"-"`

	CreationDate int    `json:"creationDate"`
	FirstAlbum   string `json:"firstAlbum"`

	Locations string `json:"-"`
	Dates     string `json:"-"`
	Relations string `json:"-"`
}

type Artists []Artist

func ParseJson() Artists {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var artists Artists

	// Read and put to struct
	err = json.Unmarshal(b, &artists)
	if err != nil {
		panic(err)
	}

	for i := range artists {
		artists[i].JoinedMembers = strings.Join(artists[i].Members, ", ")
	}

	return artists
}

func AddLocations(artists Artists) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	type LocationEntry struct {
		Locations []string `json:"locations"`
	}

	var data struct {
		Index []LocationEntry `json:"index"`
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		panic(err)
	}

	for i := range artists {
		locations := data.Index[i].Locations
		artists[i].Locations = strings.Join(locations, ", ")
	}
}

func AddRelations(artists Artists) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	type RelationsEntry struct {
		DatesLocations map[string][]string `json:"datesLocations"`
	}

	var data struct {
		Index []RelationsEntry `json:"index"`
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		panic(err)
	}

	for i := range artists {
		var s string
		for city, dates := range data.Index[i].DatesLocations {
			d := strings.Join(dates, ", ")
			s += city + ": " + d + "\n\n"
		}
		artists[i].Relations = s
	}
}

func AddDates(artists Artists) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	type DatesEntry struct {
		Dates []string `json:"dates"`
	}

	var data struct {
		Index []DatesEntry `json:"index"`
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		panic(err)
	}

	for i := range artists {
		Dates := data.Index[i].Dates
		artists[i].Dates = strings.Join(Dates, ", ")
	}
}
