package main

import (
	"fmt"
	"net/http"
	"portfolio/handlers"
)

func main() {
	fmt.Println("Starting Server at port: 8080")
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/services", handlers.Services)
	http.HandleFunc("/contact", handlers.Contact)
	http.ListenAndServe(":8080", nil)
}
