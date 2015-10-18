package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	port := flag.Int("port", 8080, "Server port")
	addr := fmt.Sprintf("localhost:%d", *port)

	r := mux.NewRouter()

	// routes
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/signup", signupHandler)

	// static resources
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))
	http.Handle("/", r)

	log.Println("Server is running on", addr)
	err := http.ListenAndServe(addr, nil)
	log.Println(err.Error())
}
