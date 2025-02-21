package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var (
	templates *template.Template
)

func init() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	fmt.Println("Initializing Game Server")

	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := struct{}{}
		templates.ExecuteTemplate(w, "index.html", data)
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server starting on http://localhost:8080")

	log.Fatal(server.ListenAndServe())
}
