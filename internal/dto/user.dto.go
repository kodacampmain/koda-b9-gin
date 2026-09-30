package dto

import "mime/multipart"

type User struct {
	// property datatype struct_tag
	Name string `json:"nama" form:"nama" example:"given"`
	Age  int8   `json:"umur" form:"umur" example:"20"`
}

type UserAccount struct {
	Email    string
	Password string
}

type EditUser struct {
	Name  string                `form:"name"`
	Image *multipart.FileHeader `form:"image"`
}
