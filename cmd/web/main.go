package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address port")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Lshortfile|log.Ltime|log.Ldate)
	app := &application{
		errorLog: errLog,
		infoLog:  infoLog,
	}
	mux := http.NewServeMux()
	// File Server
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// routing
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	// web server
	infoLog.Printf("Starting server on %v", *addr)
	err := http.ListenAndServe(*addr, mux)
	errLog.Fatal(err)
}
