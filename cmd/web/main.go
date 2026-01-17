package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address port")
	flag.Parse()
	mux := http.NewServeMux()
	// File Server
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// routing
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// web server
	log.Printf("Starting server on %v", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
