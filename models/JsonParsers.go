package models

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`

	Locations    []string `json:"-"`
	ConcertDates []string `json:"-"`
	Relations    string   `json:"-"`
}

type Artists []Artist

func ParseJson() Artists {
	// Read json
	data, err := os.ReadFile("data/artists.json")
	if err != nil {
		fmt.Println("Failed to parse:", err)
		return nil
	}

	var artists Artists

	// Read and put to struct
	err = json.Unmarshal(data, &artists)
	if err != nil {
		fmt.Println("Failed to parse:", err)
		return nil
	}

	return artists
}

func AddLocations(artists Artists) {
	b, _ := os.ReadFile("data/locations.json")

	type LocationEntry struct {
		Locations []string `json:"locations"`
	}

	var data struct {
		Index []LocationEntry `json:"index"`
	}

	err := json.Unmarshal(b, &data)
	if err != nil {
		panic(err)
	}

	for i := range artists {
		artists[i].Locations = data.Index[i].Locations
	}
}

func AddConcertDates(artists Artists) {
	b, _ := os.ReadFile("data/dates.json")

	type DatesEntry struct {
		Dates []string `json:"dates"`
	}

	var data struct {
		Index []DatesEntry `json:"index"`
	}

	err := json.Unmarshal(b, &data)
	if err != nil {
		panic(err)
	}

	for i := range artists {
		artists[i].ConcertDates = data.Index[i].Dates
	}
}

func AddRelations(artists Artists) {
	b, _ := os.ReadFile("data/relation.json")

	type RelationsEntry struct {
		DatesLocations map[string][]string `json:"datesLocations"`
	}

	var data struct {
		Index []RelationsEntry `json:"index"`
	}

	err := json.Unmarshal(b, &data)
	if err != nil {
		panic(err)
	}

	for i := range artists {
		var sb strings.Builder
		for city, dates := range data.Index[0].DatesLocations {
			d := strings.Join(dates, ", ")
			fmt.Fprintf(&sb, "%s => %s\n", city, d)
		}
		artists[i].Relations = sb.String()
	}
}
