package dal

import (
	"api/golang-song-api/confi"
	"api/golang-song-api/data_model"
	"api/golang-song-api/helpers"
	"database/sql"
	"fmt"
	"strconv"
	_ "github.com/mattn/go-sqlite3"
)

func GetAllSongs() []data_model.Song {
	var songs []data_model.Song
	var s data_model.Song
	var comments sql.NullString
	var createdBy sql.NullString
	
	db, err := sql.Open(confi.DataBaseType, confi.DataBasePath)
	checkErr(err)

	rows, err := db.Query("SELECT Id, Title, Artist, Year, Comments, CreatedBy FROM song")
	checkErr(err)

	for rows.Next() {
		err = rows.Scan(&s.Id, &s.Title, &s.Artist, &s.Year, &comments, &createdBy)
		checkErr(err)

		if comments.Valid {
			s.Comments = comments.String
		}
		if createdBy.Valid {
			s.CreatedBy = createdBy.String
		}

		songs = append(songs, s)
	}

	rows.Close()
	db.Close()

	return songs
}

func GetSong(id string) data_model.Song {
	var s data_model.Song
	var found bool
	
	db, err := sql.Open(confi.DataBaseType, confi.DataBasePath)
	checkErr(err)

	rows, err := db.Query("SELECT * FROM song WHERE ID = " + id)
	checkErr(err)

	for rows.Next() {
		err = rows.Scan(&s.Id, &s.Title, &s.Artist, &s.Year)
		checkErr(err)
		found = true
	}

	rows.Close()
	db.Close()

	if (!found) {
		panic(id + " has not been found.")
	}

	return s
}

func AddSong(s data_model.AddSong) {

	db, err := sql.Open(confi.DataBaseType, confi.DataBasePath)
	checkErr(err)

	var year int
	if s.Year != "" {
		year, _ = strconv.Atoi(s.Year)
	} 

	query := fmt.Sprintf("INSERT INTO song (Title, Artist, Year) Values ('%s', '%s', %d)", s.Title, s.Artist, year)

	_, err = db.Exec(query)
	checkErr(err)

	db.Close()
}

func DeleteSong(id string) bool {

	db, err := sql.Open(confi.DataBaseType, confi.DataBasePath)
	checkErr(err)

	query := fmt.Sprintf("DELETE FROM song WHERE ID = %s", id)

	//res, err := db.Exec("DELETE FROM song WHERE ID=$1", id)
	res, err := db.Exec(query)
	checkErr(err)

	count, err := res.RowsAffected() 
	checkErr(err)

	db.Close()

	if count > 0 {
		return true
	} else {
		return false
	}
}

func UserAuthenticate( userName string, password string) data_model.User {

	var u data_model.User 
	found := false

	db, err := sql.Open(confi.DataBaseType, confi.DataBasePath)
	checkErr(err)

	query := fmt.Sprintf("SELECT UserName, FirstName, LastName, Email, IsAdmin FROM Users WHERE UserName = '%s' AND Password = '%s'", userName, password);
	println(query + "\n")

	rows, err := db.Query(query)
	
	// Let the server send the error to the user
	if err != nil {
		checkErrNice(err)
		return u
	}	

	for rows.Next() {
		println("found some rows for the user")
		err = rows.Scan(&u.UserName, &u.FirstName, &u.LastName, &u.Email, &u.IsAdmin)

		checkErrDb(err, db)

		token := helpers.GenerateSecureToken()
		tokenExpires := helpers.GetNewTokenExpirationDate()

		u.Token = token
		u.TokenExpires = tokenExpires

		found = true
		break
	}

	rows.Close()
	db.Close()
	
	if (!found) {
		println("Authentication for user " + userName + " has failed.")
	} else {
		UpdateTokenForUser(u)
	}

	return u	
}

func UpdateTokenForUser(u data_model.User) {
	db, err := sql.Open(confi.DataBaseType, confi.DataBasePath)
	checkErr(err)

	query := fmt.Sprintf("UPDATE Users SET Token = '%s', TokenExpires = '%s' WHERE  UserName = '%s'", u.Token, u.TokenExpires.Format("01-02-2006 15:04:05"), u.UserName )
	println(query + "\n")
	_, err = db.Exec(query)
	db.Close()
	checkErrNice(err)		
}

func checkErrDb(err error, db *sql.DB) {
	db.Close()
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		println("There has been an error in the Data Access Layer")
		panic(err)
	}
}

func checkErrNice(err error) {
	if err != nil {
		println("There has been a \"nice\" error in the Data Access Layer " + err.Error())
	}
}
