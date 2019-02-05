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
	"reply.html",
)

func main() {
	http.HandleFunc("/index", index) // setting router rule
	http.HandleFunc("/login", login)
	http.HandleFunc("/reply", reply)

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

func reply(w http.ResponseWriter, r *http.Request) {

	data := struct {
		Title  string
		Header string
		Name   string
	}{
		Title:  "Reply Page",
		Header: "This is a reply",
	}

	fmt.Println("Reply")

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	data.Name = r.FormValue("username")
	data.Header = r.FormValue("password")
	fmt.Printf("Hello, %s!", data.Name)

	if err := t.ExecuteTemplate(w, "reply.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
