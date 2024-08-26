package artist

type Artist struct {
    ID           int      `json:"id"`
    Name         string   `json:"name"`
    Image        string   `json:"image"`
    Members      []string `json:"members"`
    CreationDate int      `json:"creationDate"`
    FirstAlbum   string   `json:"firstAlbum"`
    Relation     Relation `json:"relation"`
}

type Relation struct {
    ID            int                 `json:"id"`
    DatesLocations map[string][]string `json:"datesLocations"`
}
