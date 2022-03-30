package main

import(
	"net/http"
)

type Page struct {
	Data interface{}
}

type ProfilePage struct {
	User User
	ok bool
}

func GetUsers (w http.ResponseWriter, r *http.Request){
	t := template.Must(template/parseFile("templates/users.html"))

	p := Page { Data: Users, }

	t.Execute(w, p)
}
