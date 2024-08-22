package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relation"`
}

// relations api
// the whole data

// struct to represent the dates and locations
type DatesLocations map[string][]string

type RootsRelation struct {
	Index []DatesLocation
}

type DatesLocation struct {
	ArtistId int            `json:"id"`
	Places   DatesLocations `json:"datesLocations"`
}

////////////////////////////////////////////////////////////
// dates api
type RootDates struct {
	Tdates []Date `json:"index"`
}
type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

//

func fetchArtists(url string) ([]Artist, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists []Artist
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return nil, err
	}

	return artists, nil
}

func fetchDatesAndConcerts(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return
	}
	defer resp.Body.Close()

	var data RootsRelation

	// if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
	// 	return nil, err
	// }

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return
	}

	// Unmarshal the JSON data into Go structs
	// var data Data
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	// fmt.Printf("%+v\n", data)

    for _,id := range data.Index{
        println(id.ArtistId)
        for key, char := range id.Places{
            println(key,char[0])
        }
    }
}

// func handler(w http.ResponseWriter, r *http.Request) {
//     artists, err := fetchArtists("https://groupietrackers.herokuapp.com/api/artists")
//     if err != nil {
//         http.Error(w, "Secure connection failed", http.StatusInternalServerError)
//         log.Println(err)
//         return
//     }

//     tmpl := template.Must(template.ParseFiles("templates/index.html"))
//     if err := tmpl.Execute(w, artists); err != nil {
//         http.Error(w, "Loading template failed", http.StatusInternalServerError)
//         log.Println(err)
//     }
// }

func main() {
	// http.HandleFunc("/", handler)

	// staticDir := http.Dir("static")
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(staticDir)))

	// log.Print("Starting server at http://localhost:8080")
	// log.Fatal(http.ListenAndServe(":8080", nil))

	fetchDatesAndConcerts("https://groupietrackers.herokuapp.com/api/relation")


	artists, err := fetchArtists("https://groupietrackers.herokuapp.com/api/artist")
	if err != nil {
		println(err)
		return
	}

	fmt.Printf("%+v\n", artists)
}
