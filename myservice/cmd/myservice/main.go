package main

import (
	"log"
	"net/http"
)

const addr = ":8080"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	log.Println("starting server...")
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("failed to listen and serve: %v", err)
	}
}
