package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type Artist struct {
	Id           int      `json:"id"`
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

// type Data struct {
// 	Image string
// 	Members []string
// 	FirstAlbum string
// 	DatesLocation DatesLocations 
// }

// //////////////////////////////////////////////////////////
// dates api
type RootDates struct {
	Tdates []Date `json:"index"`
}
type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

//

func fetchArtists() ([]Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists []Artist
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return nil, err
	}

	return artists[:12], nil
}

func fetchDatesAndConcerts(id string) (DatesLocations, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return nil, err
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
		return nil, err
	}

	// Unmarshal the JSON data into Go structs
	// var data Data
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	var datesLocations DatesLocations

	for _, Artistid := range data.Index {
		idNum := strconv.Itoa(Artistid.ArtistId)
		if idNum == id {
			datesLocations = Artistid.Places
		}
	}

	// fmt.Printf("%+v\n", data)
	return datesLocations, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	artists, err := fetchArtists()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Println("Template 1 parsing error:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// buffer is written to memory

		// buf := &bytes.Buffer{}

		err = tmpl.Execute(w, artists)
		if err != nil {
			if err != http.ErrHandlerTimeout {
				log.Println("Template 1 execution error: ", err)
			}
		}
		// after all data is acumulated in buffer, it feed into w
		// _, err = buf.WriteTo(w)
		// if err != nil {
		// 	log.Println("Error writing response: ", err)
		// }
	}
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
	// println("jik")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")

	// println(id)

	data, err := fetchDatesAndConcerts(id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// shouldn't call http.ERror() after tmpl.Execution as headers
	// have already been written

	// fetch artists details
	tmpl, err := template.ParseFiles("templates/artistPage.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Template 2 parsing error: ", err)
		return

	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Template 2 execution error: ", err)
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/artist", artistHandler)
	// serve the static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Print("Starting server at http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// fetchDatesAndConcerts("https://groupietrackers.herokuapp.com/api/relation")

// artists, err := fetchArtists("https://groupietrackers.herokuapp.com/api/artist")
// if err != nil {
// 	println(err)
// 	return
// }

// fmt.Printf("%+v\n", artists)
