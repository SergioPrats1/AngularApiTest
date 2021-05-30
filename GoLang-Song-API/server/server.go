package server

import (
	"encoding/json"
	"net/http"
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

	a.router = r
	return a
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