package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "html/template"
)

func Any(w http.ResponseWriter, r *http.Request) {

  t := template.Must(template.ParseFiles("index.html"))

  er := "I am Muhammad"

  t.Execute(w, er)

}

func main () {

  r1 := mux.NewRouter()

  s1 := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/")))

  r1.PathPrefix("/assets/").Handler(s1)


  r1.HandleFunc("/", Any)
  http.Handle("/", r1)
 
  http.ListenAndServe(":8008", nil)

}