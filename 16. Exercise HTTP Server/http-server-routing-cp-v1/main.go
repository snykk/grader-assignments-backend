package main

import (
	"fmt"
	"net/http"
	"time"
)

func TimeHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("%s, %d %s %d", now.Weekday(), now.Day(), now.Month(), now.Year())))
	}) // TODO: replace this
}

func SayHelloHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if name == "" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello there"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
	}) // TODO: replace this
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	// TODO: answer here

	mux.HandleFunc("/time", TimeHandler())
	mux.HandleFunc("/helo", SayHelloHandler())
	return mux
}

func main() {
	// TODO: answer here
	fmt.Println("Server running on localhost:8080")

	http.ListenAndServe("localhost:8080", GetMux())
}
