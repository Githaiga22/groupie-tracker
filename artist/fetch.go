package artist

import (
    "encoding/json"
    "fmt"
    "net/http"
)

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
