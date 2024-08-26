package artist

import (
    "html/template"
    "net/http"
    "strconv"
	"log"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
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

    tmpl := template.Must(template.ParseFiles("./templates/artist.html"))
    tmpl.Execute(w, artist)
}

// handler handles requests to the main page listing all artists.
func Handler(w http.ResponseWriter, r *http.Request) {
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

