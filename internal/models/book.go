package models

// Books struct defines the structure for book data (In memory for PoC)
type Books struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
