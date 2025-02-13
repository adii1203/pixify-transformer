package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/h2non/bimg"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from EC2!")
}

func bimgVersion(w http.ResponseWriter, r *http.Request) {
	v := bimg.Version
	fmt.Fprintf(w, "bimg version: %s", v)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Default port if not set
	}

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/bimg", bimgVersion)

	log.Printf("Starting server on port %s...", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Could not start server: %s", err)
	}
}
