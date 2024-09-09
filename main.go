package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

// Artist struct to hold the artist data
type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

// Function to fetch artist data from API
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

// Handler function for the homepage
func handler(w http.ResponseWriter, r *http.Request) {
	artists, err := fetchArtists("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError, "Internal Server Error", "Secure connection failed")
		log.Println(err)
		return
	}

	if r.Method != http.MethodGet {

		renderErrorPage(w, http.StatusMethodNotAllowed, "Method", "Method not allowed")
		log.Println(err)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	if err := tmpl.Execute(w, artists); err != nil {
		renderErrorPage(w, http.StatusInternalServerError, "Internal Server Error", "Loading template failed")
		log.Println(err)
	}
}

// Function to render error pages
func renderErrorPage(w http.ResponseWriter, statusCode int, title, message string) {
	w.WriteHeader(statusCode)
	tmpl := template.Must(template.ParseFiles("templates/error.html"))
	data := struct {
		Title   string
		Message string
	}{
		Title:   title,
		Message: message,
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Failed to render error page", http.StatusInternalServerError)
	}
}

// NotFoundHandler to handle 404 errors
// func notFoundHandler(w http.ResponseWriter, r *http.Request) {
// 	renderErrorPage(w, http.StatusNotFound, "404 Not Found", "The page you are looking for does not exist.")
// }

func main() {
	// Main route
	http.HandleFunc("/", handler)

	// Static directory
	staticDir := http.Dir("static")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(staticDir)))

	// Log and start server
	log.Print("Starting server at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
