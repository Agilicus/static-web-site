package main

import (
	"log"
	"net/http"
	"os"

        _ "git.agilicus.com/open-source/static-web-site/statik"
	"github.com/gorilla/handlers"
	"github.com/rakyll/statik/fs"
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	// Serve the contents over HTTP.
	http.Handle("/", http.StripPrefix("/", http.FileServer(statikFS)))
        handler := handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)
        handler = handlers.ProxyHeaders(handler)
	http.ListenAndServe(":8080", handler)
}
