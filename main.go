package main

import (
	"html/template"
	"log"

	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "index.html")
}
func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {

	http.HandleFunc("/", homepage)

	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}
