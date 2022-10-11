package middleware

import (
	"a21hc3NpZ25tZW50/helper"
	"net/http"
)

func Get(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			helper.GenerateResponseError(w, r, 405, "Method is not allowed!")
			return
		}

		next.ServeHTTP(w, r)

	}) // TODO: replace this
}

func Post(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			helper.GenerateResponseError(w, r, 405, "Method is not allowed!")
			return
		}

		next.ServeHTTP(w, r)
	}) // TODO: replace this
}

func Delete(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			helper.GenerateResponseError(w, r, 405, "Method is not allowed!")
			return
		}

		next.ServeHTTP(w, r)
	}) // TODO: replace this
}
