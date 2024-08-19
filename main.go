package main

import (
    "encoding/json"
    "net/http"
    "log"
    "html/template"
)

type Artist struct {
    ID          int      `json:"id"`
    Name        string   `json:"name"`
    Image       string   `json:"image"`
    Members     []string `json:"members"`
    CreationDate int     `json:"creationDate"`
    FirstAlbum  string   `json:"firstAlbum"`
}

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

    return artists, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
    artists, err := fetchArtists()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    tmpl := template.Must(template.ParseFiles("templates/index.html"))
    if err := tmpl.Execute(w, artists); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main() {
    http.HandleFunc("/", handler)
	log.Print("Starting server at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
