package middleware

import (
	"net/http"
	"time"

	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/helper"
	"a21hc3NpZ25tZW50/model"
)

func isExpired(s model.Session) bool {
	return s.Expiry.Before(time.Now())
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		coockie, err := r.Cookie("session_token")

		if err != nil {
			helper.GenerateResponseError(w, r, 401, "http: named cookie not present")
			return
		}

		for key := range db.Sessions {
			if key == coockie.Value && !isExpired(db.Sessions[key]) {
				next.ServeHTTP(w, r)
				return
			}
		}

		helper.GenerateResponseError(w, r, 401, "http: named cookie not present")
	}) // TODO: replace this
}
