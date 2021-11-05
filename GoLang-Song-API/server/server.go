package server

import (
	"api/golang-song-api/dal"
	"api/golang-song-api/data_model"
	"api/golang-song-api/token"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
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
	r.HandleFunc("/getSong/{ID:[a-zA-Z0-9_]+}", a.fetchSong).Methods(http.MethodGet, http.MethodOptions)

	r.HandleFunc("/addSong", a.addSong).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/deleteSong/{ID:[a-zA-Z0-9_]+}", a.deleteSong).Methods(http.MethodGet, http.MethodOptions)

	r.HandleFunc("/retrieveParameterTest", a.retrieveParameterTest).Methods(http.MethodGet, http.MethodOptions)

	r.HandleFunc("/users/authenticate", a.userAuthenticate).Methods(http.MethodPost, http.MethodOptions)

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) deleteSong(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	data, err := token_manager.ValidateJwcToken(r)

    if err != nil {
        println(err)
        http.Error(w, "Request failed!", http.StatusUnauthorized)
		return
    }

	if data == nil {
		return
	}

	userName := data.CustomClaims["userName"]

	vars := mux.Vars(r)
	id := vars["ID"]

	result := dal.DeleteSong(id, userName)

	w.Header().Set("Content-Type", "application/json")
	if result {
		json.NewEncoder(w).Encode("The song whose id is " + id + " has been deleted added")
	} else {
		json.NewEncoder(w).Encode("The song whose id is " + id + " has not been found")
	}
}

func (a *api) addSong(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	data, err := token_manager.ValidateJwcToken(r)

    if err != nil {
        println(err)
        http.Error(w, "Request failed!", http.StatusUnauthorized)
		return
    }

	if data == nil {
		return
	}

	userName := data.CustomClaims["userName"]
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

	songToAdd.CreatedBy = userName

	dal.AddSong(songToAdd)

	//w.Header().Set("Content-Type", "text/html")
	//w.Write( []byte("The song " + songToAdd.Artist + " has been added") )

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("The song " + songToAdd.Artist + " has been added")
}


func (a *api) fetchSongs(w http.ResponseWriter, r *http.Request) {
	var songList []data_model.Song

	enableCors(&w)

	//GetHeaders(r)

	data, err := token_manager.ValidateJwcToken(r)

    if err != nil {
        println(err)
        http.Error(w, "Request failed!", http.StatusUnauthorized)
		return
    }

	if (data == nil) {
		return
	}

    userName := data.CustomClaims["userName"]
	
	songList = dal.GetAllSongs(userName)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(songList)
}

// This function is not called from the Angular App
func (a *api) fetchSong(w http.ResponseWriter, r *http.Request) {
	var song  data_model.Song
	
	enableCors(&w)

	vars := mux.Vars(r)
	id := vars["ID"]

	song = dal.GetSong(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(song)
}

func (a *api) userAuthenticate(w http.ResponseWriter, r *http.Request) {

	// This needs to be called before decoding the payload!
	enableCors(&w)

	if (r.Method == http.MethodOptions) {
		return;
	}

	var u data_model.UserPassword;

	err := json.NewDecoder(r.Body).Decode(&u)

    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	var user data_model.User

	user = dal.UserAuthenticate(u.UserName, u.Password);

	if user == (data_model.User{}) {
        http.Error(w, "Authentication failed for user " + u.UserName, http.StatusBadRequest)
        return
    }

	user.Token = token_manager.GeneratJwcToken(user.UserName)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}


func (a *api) retrieveParameterTest(w http.ResponseWriter, r *http.Request) {
	
	keys, ok := r.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
        panic("Url Param 'key' is missing")
    }

    // Query()["key"] will return an array of items, 
    // we only want the single item.
    key := keys[0]

	enableCors(&w)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("The key that was passed is: " + key)
}


func GetHeaders(r *http.Request) {
	for name, values := range r.Header {
		// Loop over all values for the name.
		for _, value := range values {
			fmt.Println(name, " : ", value)
		}
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
    (*w).Header().Set("Access-Control-Allow-Headers", "*")
}