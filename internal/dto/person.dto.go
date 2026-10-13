package dto

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type PersonQuery struct {
	Limit int `form:"limit"`
	Page  int `form:"page"`
}
