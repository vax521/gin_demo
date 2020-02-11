package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.FileServer(http.Dir("./static_demo")))
	server := &http.Server{
		Addr:    ":8090",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
