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

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
	fileserver := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)
	addr := flag.String("addr", ":4000", "Http Network Address: ")
	flag.Parse()

	infoLog.Printf("Starting Server on port: %s", *addr)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog, Handler: mux,
	}

	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
