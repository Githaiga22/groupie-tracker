package main

import (
    "encoding/json"
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strconv"
)

// Artist represents the structure of an artist's data.
type Artist struct {
    ID           int      `json:"id"`
    Name         string   `json:"name"`
    Image        string   `json:"image"`
    Members      []string `json:"members"`
    CreationDate int      `json:"creationDate"`
    FirstAlbum   string   `json:"firstAlbum"`
    Relation     Relation `json:"relation"`
}

// Relation represents the structure for events related to an artist.
type Relation struct {
    ID            int                 `json:"id"`
    DatesLocations map[string][]string `json:"datesLocations"`
}

// fetchArtistData fetches data for all artists from the given URL.
func fetchArtistData(url string) ([]Artist, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var artists []Artist
    if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
        return nil, err
    }

    for i, artist := range artists {
        relationUrl := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%d", artist.ID)
        relationResp, err := http.Get(relationUrl)
        if err != nil {
            return nil, err
        }
        defer relationResp.Body.Close()

        var relation Relation
        if err := json.NewDecoder(relationResp.Body).Decode(&relation); err != nil {
            return nil, err
        }

        artists[i].Relation = relation
    }

    return artists, nil
}

// fetchArtists fetches the list of all artists.
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

// getArtistByID returns a single artist by ID.
func getArtistByID(id int) (Artist, error) {
    artists, err := fetchArtistData("https://groupietrackers.herokuapp.com/api/artists")
    if err != nil {
        return Artist{}, err
    }

    for _, artist := range artists {
        if artist.ID == id {
            return artist, nil
        }
    }

    return Artist{}, fmt.Errorf("artist not found")
}

// artistHandler handles requests to the artist detail page.
func artistHandler(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Path[len("/artist/"):]
    artistID, err := strconv.Atoi(id)
    if err != nil {
        http.Error(w, "Invalid artist ID", http.StatusBadRequest)
        return
    }

    artist, err := getArtistByID(artistID)
    if err != nil {
        http.Error(w, "Artist not found", http.StatusNotFound)
        return
    }

    tmpl := template.Must(template.ParseFiles("templates/artist.html"))
    tmpl.Execute(w, artist)
}

// handler handles requests to the main page listing all artists.
func handler(w http.ResponseWriter, r *http.Request) {
    artists, err := fetchArtists("https://groupietrackers.herokuapp.com/api/artists")
    if err != nil {
        http.Error(w, "Failed to load artists", http.StatusInternalServerError)
        log.Println(err)
        return
    }

    tmpl := template.Must(template.ParseFiles("templates/index.html"))
    if err := tmpl.Execute(w, artists); err != nil {
        http.Error(w, "Failed to load template", http.StatusInternalServerError)
        log.Println(err)
    }
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/artist/", artistHandler)

    staticDir := http.Dir("static")
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(staticDir)))

    log.Print("Starting server at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
