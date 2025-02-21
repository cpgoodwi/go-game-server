package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Initializing Game Server")

	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server starting on port 8080")

	log.Fatal(server.ListenAndServe())
}
