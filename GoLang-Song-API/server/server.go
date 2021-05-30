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
	r.HandleFunc("/getSongs", a.fetchSongs).Methods(http.MethodGet)
	r.HandleFunc("/getSongs/{ID:[a-zA-Z0-9_]+}", a.fetchSong).Methods(http.MethodGet)

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) fetchSongs(w http.ResponseWriter, r *http.Request) {
	var songList []data_model.Song
	songList = dal.GetAllSongs()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(songList)
}

func (a *api) fetchSong(w http.ResponseWriter, r *http.Request) {
	var song  data_model.Song
	
	vars := mux.Vars(r)
	id := vars["ID"]

	song = dal.GetSong(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(song)
}