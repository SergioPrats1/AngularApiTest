package server

import (
	"api/golang-song-api/dal"
	"api/golang-song-api/data_model"
	"api/golang-song-api/token"
	"encoding/json"
	"net/http"
	"io/ioutil"
)


func (a *api) userAuthenticate(w http.ResponseWriter, r *http.Request) {

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


func (a *api) userRegister(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	if r.Method == http.MethodOptions {
		return
	}

	var newUser data_model.User

	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &newUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		panic(err)
	}

	if (dal.CheckUserExists(newUser.UserName)) 	{
		http.Error(w, "The user " + newUser.UserName + " already exists" , http.StatusNotAcceptable)
		return
	}

	dal.AddUser(newUser)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("The User " + newUser.UserName + " has been added")
}