package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./template/*.html"))
}

func index(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func services(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "services.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "contact.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Starting Server at port: 8080")
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/services", services)
	http.HandleFunc("/contact", contact)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
