package dal

import (
	"api/golang-song-api/data_model"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func GetAllSongs() []data_model.Song {
	var songs []data_model.Song
	var s data_model.Song
	
	db, err := sql.Open("sqlite3", "C:\\Sqlite\\Songs.db")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM song")
	checkErr(err)

	for rows.Next() {
		err = rows.Scan(&s.Id, &s.Title, &s.Artist, &s.Year)
		checkErr(err)
		songs = append(songs, s)
	}

	rows.Close()
	db.Close()

	return songs
}

func GetSong(id string) data_model.Song {
	var s data_model.Song
	var found bool
	//id_str := strconv.Itoa(id)
	
	db, err := sql.Open("sqlite3", "C:\\Sqlite\\Songs.db")
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


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}