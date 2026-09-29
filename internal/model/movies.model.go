package model

type Movies struct {
	Title string `db:"title"`
}

type MoviesWithGenres struct {
	Title  string `db:"title"`
	Genres string `db:"genres"`
}
