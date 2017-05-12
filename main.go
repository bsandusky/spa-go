package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("static/templates/*.gohtml"))
}

func main() {
	mux := http.NewServeMux()

	// Front end server
	mux.HandleFunc("/", serveTemplateHandler)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	mux.Handle("/favicon.ico", http.NotFoundHandler())

	// Backend APIs
	// TKTKTKTK

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func serveTemplateHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "base.gohtml", "The Server")
}
