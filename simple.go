package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server at localhost :8000...")
	http.HandleFunc("/ping", logger(Ping))
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func logger(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s - %s", r.RemoteAddr, r.RequestURI, r.Method)
		f(w, r)
	}
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Pong\n")
}
