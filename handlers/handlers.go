package handlers

import (
    "encoding/json"
    "html/template"
    "log"
    "net/http"
)

type Artist struct {
    ID           int      `json:"id"`
    Name         string   `json:"name"`
    Image        string   `json:"image"`
    Members      []string `json:"members"`
    CreationDate int      `json:"creationDate"`
    FirstAlbum   string   `json:"firstAlbum"`
}

// Fetch artists from the API
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

// Homepage handler
func HomepageHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/homepage.html"))
    if err := tmpl.Execute(w, nil); err != nil {
        http.Error(w, "Loading template failed", http.StatusInternalServerError)
        log.Println(err)
    }
}

// Artists handler
func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
    artists, err := fetchArtists("https://groupietrackers.herokuapp.com/api/artists")
    if err != nil {
        http.Error(w, "Failed to fetch artists", http.StatusInternalServerError)
        log.Println(err)
        return
    }

    tmpl := template.Must(template.ParseFiles("templates/index.html"))
    if err := tmpl.Execute(w, artists); err != nil {
        http.Error(w, "Loading template failed", http.StatusInternalServerError)
        log.Println(err)
    }
}

// Add other handlers for locations, dates, and relations similarly...
