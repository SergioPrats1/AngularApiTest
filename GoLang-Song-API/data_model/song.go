package data_model

type Song struct {
	Id     int      `json:"id"`
	Title  string	`json:"title,omitempty"`
	Artist string	`json:"artist,omitempty"`
	Year   int		`json:"year,omitempty"`
}
