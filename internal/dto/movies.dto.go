package dto

type AddMovies struct {
	Title  string `json:"title"`
	Genres []int  `json:"genres"`
}
