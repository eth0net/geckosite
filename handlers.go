package main

import (
	"html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	lp, hp := "templates/layout.gohtml", "templates/home.gohtml"
	tmpl := template.Must(template.ParseFiles(lp, hp))
	tmpl.ExecuteTemplate(w, "layout", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	lp, hp := "templates/layout.gohtml", "templates/about.gohtml"
	tmpl := template.Must(template.ParseFiles(lp, hp))
	tmpl.ExecuteTemplate(w, "layout", nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	lp, hp := "templates/layout.gohtml", "templates/contact.gohtml"
	tmpl := template.Must(template.ParseFiles(lp, hp))
	tmpl.ExecuteTemplate(w, "layout", nil)
}
