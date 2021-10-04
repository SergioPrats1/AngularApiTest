package data_model

import(
	"time"
)

type User struct {
	Id     int      	`json:"id"`
	UserName  string	`json:"username"`
	Password string		`json:"password"`
	FirstName string	`json:"firstname"`
	LastName string		`json:"lastname"`
	Email string		`json:"email"`
	IsAdmin bool		`json:"isadmin"`
	Token string		`json:"token"`
	TokenExpires time.Time	`json:"tokenexpires"`
}

type UserPassword struct {
	UserName  string	
	Password string		
}