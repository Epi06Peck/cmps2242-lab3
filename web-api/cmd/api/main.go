package main

import (
	"log"
	"net/http"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Shapes API"))
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is running"))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Epi Peck"))
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(time.Now().String()))
}

func random(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("11"))
}
func greeting(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello user!"))
}

func quote(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Never lose hope. You never know what tomorrow may bring!"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/health", health)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/timeH", timeHandler)
	mux.HandleFunc("/greeting", greeting)
	mux.HandleFunc("/quote", quote)

	log.Print("Starting Server on port 4000")
	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)

}
