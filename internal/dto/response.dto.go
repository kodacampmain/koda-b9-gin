package dto

type Response struct {
	Success bool   `example:"true"`
	Data    any    ``
	Msg     string `example:"berhasil"`
}

type ErrorResponse struct {
	Success bool   `example:"false"`
	Data    any    ``
	Msg     string `example:"berhasil"`
}
