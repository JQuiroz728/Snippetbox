package main

import (
	"log"
	"net/http"
)

// Handler Functions
func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/". If it doesn't, use
	// the http.NotFound() function to send a 404 response to the client.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox - Jorge"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using POST or not.
	// http.MethodPost is a constant equal to the string "POST".
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed) // http.StatusMethodNotAllowed = 405 status code
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my about page"))
}

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
