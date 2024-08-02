package models

// Album representa un álbum musical
type Album struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Artist string `json:"artist"`
    Year   int    `json:"year"`
}
