package main_test

import (
	main "a21hc3NpZ25tZW50"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Chaining Middleware", func() {
	Describe("Test endpoint /admin with different method", func() {
		When("hit with POST request method and Header role is ADMIN", func() {
			It("should return method not allowed", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("POST", "/admin", nil)
				r.Header.Set("role", "ADMIN")
				handler := main.RequestMethodGetMiddleware(main.AdminMiddleware(main.AdminHandler()))
				handler.ServeHTTP(w, r)
				Expect(w.Code).To(Equal(http.StatusMethodNotAllowed))
				Expect(w.Body.String()).To(Equal("Method is not allowed"))
			})
		})

		When("hit with POST request method and no role", func() {
			It("should return method not allowed", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("POST", "/admin", nil)
				handler := main.RequestMethodGetMiddleware(main.AdminMiddleware(main.AdminHandler()))
				handler.ServeHTTP(w, r)
				Expect(w.Code).To(Equal(http.StatusMethodNotAllowed))
				Expect(w.Body.String()).To(Equal("Method is not allowed"))
			})
		})

		When("it with GET request method and Header role is not ADMIN", func() {
			It("should return unauthorized", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/admin", nil)
				r.Header.Set("role", "USER")
				handler := main.RequestMethodGetMiddleware(main.AdminMiddleware(main.AdminHandler()))
				handler.ServeHTTP(w, r)
				Expect(w.Code).To(Equal(http.StatusUnauthorized))
				Expect(w.Body.String()).To(Equal("Role not authorized"))
			})
		})

		When("hit with GET request method and no role", func() {
			It("should return unauthorized", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/admin", nil)
				handler := main.RequestMethodGetMiddleware(main.AdminMiddleware(main.AdminHandler()))
				handler.ServeHTTP(w, r)
				Expect(w.Code).To(Equal(http.StatusUnauthorized))
				Expect(w.Body.String()).To(Equal("Role not authorized"))
			})
		})

		When("hit with GET request method and Header role is ADMIN", func() {
			It("should return method not allowed", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/admin", nil)
				r.Header.Set("role", "ADMIN")
				handler := main.RequestMethodGetMiddleware(main.AdminMiddleware(main.AdminHandler()))
				handler.ServeHTTP(w, r)
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(w.Body.String()).To(Equal("Welcome to Admin page"))
			})
		})
	})
})
