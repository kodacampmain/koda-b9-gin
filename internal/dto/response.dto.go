package dto

type Response struct {
	Success  bool           `example:"true"`
	Data     any            ``
	Msg      string         `example:"berhasil"`
	Metadata PaginationMeta `json:"metadata,omitzero"`
}

type PaginationMeta struct {
	// TotalData   int
	CurrentPage int     `json:"current_page"`
	TotalPage   float64 `json:"total_page"`
	Next        string  `json:"next"`
	Prev        string  `json:"prev"`
}

type ErrorResponse struct {
	Success bool   `example:"false"`
	Data    any    ``
	Msg     string `example:"berhasil"`
}
