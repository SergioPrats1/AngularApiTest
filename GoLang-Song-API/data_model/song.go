package data_model

type Song struct {
	Id     int      `json:"ID"`
	Title  string	`json:"Title,omitempty"`
	Artist string	`json:"Artist,omitempty"`
	Year   int		`json:"Year,omitempty"`
}
