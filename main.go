package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	r.Path("/favicon.ico").Handler(fs)

	r.Path("/").HandlerFunc(homeHandler)
	r.Path("/about").HandlerFunc(aboutHandler)
	r.Path("/contact").HandlerFunc(contactHandler)

	http.ListenAndServe(":8080", r)
}
