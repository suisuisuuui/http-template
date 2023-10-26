package main

import (
	"log"
	"net/http"
)

func sampleHandler(w http.ResponseWriter, r *http.Request) {
	body := getHTTPBody(r)
	w.WriteHeader(200)
	w.Write(body)
}

// Convert http request body to []byte
func getHTTPBody(r *http.Request) []byte {
	len := r.ContentLength
	b := make([]byte, len)
	r.Body.Read(b)
	return b
}

func main() {

	// Register Handler
	http.HandleFunc("/sample", sampleHandler)

	// Start HTTP Server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("listen and serve error: %w", err)
	}
}
