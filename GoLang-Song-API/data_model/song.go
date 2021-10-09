package data_model

type Song struct {
	Id     int      `json:"id"`
	Title  string	`json:"title,omitempty"`
	Artist string	`json:"artist,omitempty"`
	Year   int		`json:"year,omitempty"`
	Comments string `json:"comments,omitempty"`
	CreatedBy string `json:"createdby,omitempty"`
}

type AddSong struct {
	Title  string	`json:"title,omitempty"`
	Artist string	`json:"artist,omitempty"`
	Year   string	`json:"year,omitempty"`
	Comments string `json:"comments,omitempty"`
	CreatedBy string `json:"createdby,omitempty"`	
   }