package server

import (
	"encoding/json"
	"net/http"
	"fmt"
	"io/ioutil"
	"strconv"
	"github.com/gorilla/mux"
	"api/golang-song-api/data_model"
	"api/golang-song-api/dal"
)

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

func New() Server {
	a := &api{}

	r := mux.NewRouter()
	r.HandleFunc("/getSongs", a.fetchSongs).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/getSongs/{ID:[a-zA-Z0-9_]+}", a.fetchSong).Methods(http.MethodGet, http.MethodOptions)

	r.HandleFunc("/addSong", a.addSong).Methods(http.MethodPost, http.MethodOptions)

	a.router = r
	return a
}

func (a *api) addSong(w http.ResponseWriter, r *http.Request) {

	var songToAdd data_model.AddSong
	
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &songToAdd)

	if err != nil {
		fmt.Println("into a panic statement in server.addSong()")
		panic(err)
	}

	if (len(songToAdd.Title) == 0 || len(songToAdd.Artist) == 0) {
		panic ("Songs cannot be added to database if the Artist or Title fields are empty!")
	}

	if songToAdd.Year != "" {
		if  _, err := strconv.Atoi(songToAdd.Year); err != nil {
			panic("The song's year is not a number")
		}
	}

	dal.AddSong(songToAdd)

	enableCors(&w)

	//w.Header().Set("Content-Type", "text/html")
	//w.Write( []byte("The song " + songToAdd.Artist + " has been added") )

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("The song " + songToAdd.Artist + " has been added")
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) fetchSongs(w http.ResponseWriter, r *http.Request) {
	var songList []data_model.Song
	songList = dal.GetAllSongs()

	enableCors(&w)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(songList)
}

func (a *api) fetchSong(w http.ResponseWriter, r *http.Request) {
	var song  data_model.Song
	
	vars := mux.Vars(r)
	id := vars["ID"]

	song = dal.GetSong(id)

	enableCors(&w)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(song)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
    (*w).Header().Set("Access-Control-Allow-Headers", "*")
}