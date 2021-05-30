package main

import (
	"api/golang-song-api/server"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Running...")

	serv := server.New()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", serv.Router()))
}

func handler(w http.ResponseWriter, r *http.Request) { }

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}