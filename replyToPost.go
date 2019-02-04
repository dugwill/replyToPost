package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var t, _ = template.ParseFiles(
	"index.html",
)

func main() {
	http.HandleFunc("/index", index) // setting router rule

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

	if err := t.ExecuteTemplate(w, "newEventHandle.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//t, _ := template.ParseFiles("index.html")
	//t.ExecuteTemplate(w, "index.html", data)

}
