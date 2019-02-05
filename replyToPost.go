package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var t, _ = template.ParseFiles(
	"index.html",
	"login.html",
	"bob.html",
)

func main() {
	http.HandleFunc("/index", index) // setting router rule
	http.HandleFunc("/login", login)
	http.HandleFunc("/bob", bob)

	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {

	data := struct {
		Title  string
		Header string
		Name   string
	}{
		Title:  "Index page",
		Header: "Index",
		Name:   "Index Page",
	}

	fmt.Println("index")

	if err := t.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {

	data := struct {
		Title  string
		Header string
		Name   string
	}{
		Title:  "Login page",
		Header: "Login",
	}

	fmt.Println("index")

	if err := t.ExecuteTemplate(w, "login.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//t, _ := template.ParseFiles("index.html")
	//t.ExecuteTemplate(w, "index.html", data)

}

func bob(w http.ResponseWriter, r *http.Request) {

	data := struct {
		Title  string
		Header string
		Name   string
	}{
		Title:  "Bob Page",
		Header: "Bob",
	}

	fmt.Println("bob")

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	data.Name = r.FormValue("username")
	data.Header = r.FormValue("password")
	fmt.Printf("Hello, %s!", data.Name)

	if err := t.ExecuteTemplate(w, "bob.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
