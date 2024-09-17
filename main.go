package main

import (
	"log"
	"net/http"

	"tracker/handlers"
)

// Function to render error pages

func main() {
	http.HandleFunc("/", handlers.HomepageHandler)
	http.HandleFunc("/artist", handlers.ArtistHandler)
	http.HandleFunc("/dates", handlers.DateHandler)
	http.HandleFunc("/locations", handlers.LocationHandler)
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
