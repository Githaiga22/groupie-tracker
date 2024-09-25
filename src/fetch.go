package src

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	helper "tracker/helpers"
)

var Data helper.Data

func FetchArtists() ([]helper.Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists []helper.Artist
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return nil, err
	}
	return artists, nil
}

func FetchLocations(id string) (helper.Location, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil  {
		fmt.Println("Error reading the response body:", err)
		return helper.Location{}, err
	}
	defer resp.Body.Close()
	
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return helper.Location{}, err
	}
	// Unmarshal the JSON data into Go structs
	var data helper.AllLocations
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return helper.Location{}, err
	}

	Data.Locations = data.Location
	var locations helper.Location	

	for _, Artistid := range data.Location {
		idNum,_ := strconv.Atoi(id)	
		if Artistid.ArtistId == idNum {
			locations = Artistid
			break
		}
	}
	return locations, nil
}

func FetchDates(id string) (helper.Date, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return helper.Date{}, err
	}
	defer resp.Body.Close()
	
	var data helper.RootDates
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return helper.Date{}, err
	}
	// Unmarshal the JSON data into Go structs
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return helper.Date{}, err
	}

	var dates helper.Date

	Data.Dates = data.Tdates

	for _, Artistid := range data.Tdates {
		idNum := strconv.Itoa(Artistid.Id)
		if idNum == id {
			dates = Artistid
		}
	}

	for i, date := range dates.Dates{
		if date[0] == '*'{
			dates.Dates[i] = date[1:]
		}
	}

	return dates, nil
}

func FetchDatesAndConcerts(id string) (helper.DatesLocations, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return nil, err
	}
	defer resp.Body.Close()

	var data helper.RootsRelation

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return nil, err
	}

	// Unmarshal the JSON data into Go structs
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	var datesLocations helper.DatesLocations

	for _, Artistid := range data.Relation {
		idNum := strconv.Itoa(Artistid.Id)
		if idNum == id {
			datesLocations = Artistid.Places
		}
	}

	return datesLocations, nil
}
