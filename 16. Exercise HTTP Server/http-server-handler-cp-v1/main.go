package main

import (
	"fmt"
	"net/http"
	"time"
)

func GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		now := time.Now()

		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(fmt.Sprintf("%s, %d %s %d", now.Weekday(), now.Day(), now.Month(), now.Year())))
	} // TODO: replace this
}

func main() {
	fmt.Println("Server running on localhost:8080")

	http.ListenAndServe("localhost:8080", GetHandler())
}
