package main

import (
    "log"
    "net/http"
    "tracker/artist"
)

func main() {
    http.HandleFunc("/", artist.Handler)
    http.HandleFunc("/artist/", artist.ArtistHandler)

    staticDir := http.Dir("static")
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(staticDir)))

    log.Print("Starting server at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
