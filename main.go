package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	s := http.StripPrefix("/x/", http.FileServer(http.Dir("./main.css/")))

	r.PathPrefix("/x/").Handler(s)

	r.HandleFunc("/users", GetUsers)
	r.HandleFunc("/profile/{username}", GetProfile)
	http.Handle("/", r)
	http.ListenAndSere(":8080",nil)
}