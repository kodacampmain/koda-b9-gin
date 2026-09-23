package dto

type HeaderQuery struct {
	Title  string   `form:"title"`
	Genres []string `form:"genre"`
}
