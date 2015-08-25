package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("port", 8080, "Server port")
	addr := fmt.Sprintf("localhost:%d", *port)

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))
	http.Handle("/", r)

	log.Println("Server is running on", addr)
	err := http.ListenAndServe(addr, nil)
	log.Println(err.Error())
}
