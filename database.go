package main

type User struct {
	Username, FullName string
	Skills []string
}

var Users = []User {
	User {
		Username: "Muhammad"
		FullName: "MuhammadSarvar"
		Skills: []string { "Go", "Fyne", "Python", "C"}
	},
	User {
		Username: "Elyor"
		FullName: "Elyor Aka"
		Skills: []string { "Go", "Python", "C", "Linux"}
	},
	User {
		Username: "Farruh"
		FullName: "Faruh Ismoilov"
		Skills: []string { "Go", "Python", "C"}
	},



}