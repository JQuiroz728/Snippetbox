package main

import (
	"log"
	"net/http"
)

func main() {
	// Use the http.NewServeMux() function to initialize a new
	// servemux, then register the home function as the handler
	// for the "/" URL pattern/route.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	mux.HandleFunc("/about", about)

	// http.ListenAndServe() starts a new web server.
	// Two parameters: the TCP network address to listen on (in
	// this case ":4000") and the servemux that was created. If
	// http.ListenAndServer() returns an error, use the log.Fatal()
	// function to log the error message and exit. Any error returned
	// by ListenAndServe is always non-nil.
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
