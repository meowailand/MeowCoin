package main

import (
	"html/template"
	"log"
	"net/http"
)

const port string = ":4000"

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.html")

}

func main() {
	log.Fatal(http.ListenAndServe(port, nil))
}
