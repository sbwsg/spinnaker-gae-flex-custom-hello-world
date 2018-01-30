package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	config := newServerConfig()

	// app engine healthcheck endpoint
	http.HandleFunc("/_ah/health", http.HandlerFunc(healthHandler))

	// hello world endpoint
	http.HandleFunc("/", http.HandlerFunc(indexHandler))

	log.Printf("Listening on %s:%s...", config.addr, config.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", config.addr, config.port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.Write([]byte("hello, world!"))
	} else {
		http.NotFound(w, r)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
