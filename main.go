package main

import (
	"log"
	"net/http"
	"os"

	_ "git.agilicus.com/open-source/static-web-site/statik"
	"github.com/MEDIGO/go-healthz"
	"github.com/gorilla/handlers"
	"github.com/rakyll/statik/fs"
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	healthz.Set("version", "0.1")

	// Serve the contents over HTTP.
	http.Handle("/", http.StripPrefix("/", http.FileServer(statikFS)))
	http.Handle("/healthz", healthz.Handler())
	handler := handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)
	handler = handlers.ProxyHeaders(handler)
	bind := os.Getenv("STATIC_WEB_SITE_BIND")
	if bind == "" {
		bind = ":8080"
	}
	http.ListenAndServe(bind, handler)
}
