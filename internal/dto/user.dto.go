package dto

type User struct {
	// property datatype struct_tag
	Name string `json:"nama" form:"nama"`
	Age  int8   `json:"umur" form:"umur"`
}

type Account struct {
	Email    string
	Password string
}
