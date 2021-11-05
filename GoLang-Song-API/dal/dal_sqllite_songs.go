package dal

import (
	"api/golang-song-api/confi"
	"api/golang-song-api/data_model"
	"database/sql"
	"fmt"
	"strconv"
	_ "github.com/mattn/go-sqlite3"
)

func GetAllSongs(userName string) []data_model.Song {
	var songs []data_model.Song
	var s data_model.Song
	var comments sql.NullString
	var createdBy sql.NullString
	
    isAdmin, err := getUserAdmin(userName)
	checkErr(err)

	db, err := sql.Open(confi.DataBaseType, confi.DataBasePath)
	checkErr(err)

	query := "SELECT Id, Title, Artist, Year, Comments, CreatedBy FROM song"

	if (!isAdmin){
		query = query + " WHERE CreatedBy = '" + userName + "'"
	}

	rows, err := db.Query(query)
	checkErr(err)

	for rows.Next() {
		err = rows.Scan(&s.Id, &s.Title, &s.Artist, &s.Year, &comments, &createdBy)
		checkErr(err)

		if comments.Valid {
			s.Comments = comments.String
		} else{
			s.Comments = ""
		}

		if createdBy.Valid {
			s.CreatedBy = createdBy.String
		} else{
			s.CreatedBy = ""			
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

	query := fmt.Sprintf("INSERT INTO song (Title, Artist, Year, CreatedBy, Comments) Values ('%s', '%s', %d, '%s', '%s')",
			s.Title, s.Artist, year, s.CreatedBy, s.Comments)

	_, err = db.Exec(query)
	checkErr(err)

	db.Close()
}

func DeleteSong(id string, userName string) bool {

    isAdmin, err := getUserAdmin(userName)
	checkErr(err)

	db, err := sql.Open(confi.DataBaseType, confi.DataBasePath)
	checkErr(err)

	query := fmt.Sprintf("DELETE FROM song WHERE ID = %s", id)

	if(!isAdmin) {
		query = query + fmt.Sprintf( " AND CreatedBy ='%s'", userName)
	}

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
