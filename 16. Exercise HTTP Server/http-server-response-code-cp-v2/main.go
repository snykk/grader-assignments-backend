package main

import (
	"net/http"
)

var students = []string{
	"Aditira",
	"Dito",
	"Afis",
	"Eddy",
}

func IsNameExists(name string) bool {
	for _, n := range students {
		if n == name {
			return true
		}
	}

	return false
}

func CheckStudentName() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		studentName := r.URL.Query().Get("name")

		if r.Method != "GET" {
			w.WriteHeader(405)
			w.Write([]byte("Method is not allowed"))
			return
		}

		if IsNameExists(studentName) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Name is exists"))
			return
		}

		w.WriteHeader(404)
		w.Write([]byte("Data not found"))
	})
	// TODO: replace this
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	// TODO: answer here

	mux.HandleFunc("/students", CheckStudentName())
	return mux
}

func main() {
	http.ListenAndServe("localhost:8080", GetMux())
}
