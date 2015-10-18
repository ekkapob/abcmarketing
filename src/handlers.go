package main

import (
	"github.com/yosssi/ace"
	// "html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tpl, _ := ace.Load(
		"templates/base", "templates/home",
		&ace.Options{DynamicReload: true})
	// tpl.Execute(w, []string{"1", "2", "3", "4", "5", "6"})
	tpl.Execute(w, nil)

	// t, _ := template.ParseFiles("templates/home.html", "templates/_header.html")
	// t.Execute(w, "Hello World!")
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	tpl, _ := ace.Load(
		"templates/base", "templates/signup",
		&ace.Options{DynamicReload: true})
	tpl.Execute(w, nil)
}
