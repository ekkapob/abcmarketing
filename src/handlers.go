package main

import (
	"html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/home.html", "templates/header.html")
	t.Execute(w, "Hello World!")
}
