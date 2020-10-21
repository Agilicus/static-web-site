package main

import (
	"log"
	"net/http"

	_ "./statik"
	"github.com/rakyll/statik/fs"
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	// Serve the contents over HTTP.
	http.Handle("/", http.StripPrefix("/", http.FileServer(statikFS)))
	http.ListenAndServe(":8080", nil)
}
