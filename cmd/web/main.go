package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address port")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Lshortfile|log.Ltime|log.Ldate)
	mux := http.NewServeMux()
	// File Server
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// routing
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// web server
	infoLog.Printf("Starting server on %v", *addr)
	err := http.ListenAndServe(*addr, mux)
	errLog.Fatal(err)
}
