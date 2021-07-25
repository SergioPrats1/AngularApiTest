package dal

import (
	"api/golang-song-api/confi"
	"api/golang-song-api/data_model"
	"database/sql"
	"fmt"
	"strconv"
	_ "github.com/mattn/go-sqlite3"
)

func GetAllSongs() []data_model.Song {
	var songs []data_model.Song
	var s data_model.Song
	
	db, err := sql.Open(confi.DataBaseType, confi.DataBasePath)
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

	year, _ := strconv.Atoi(s.Year)

	query := fmt.Sprintf("INSERT INTO song (Title, Artist, Year) Values ('%s', '%s', %d)", s.Title, s.Artist,year)

	//fmt.Printf("query: %s \n", query)
	_, err = db.Exec(query)
	checkErr(err)

	db.Close()
}


func checkErr(err error) {
	if err != nil {
		println("There has been an error in the dal")
		panic(err)
	}
}