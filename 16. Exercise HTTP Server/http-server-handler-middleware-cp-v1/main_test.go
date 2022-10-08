package main_test

import (
	main "a21hc3NpZ25tZW50"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Middleware", func() {
	Describe("Test endpoint /student with different method", func() {
		When("hit with GET request method", func() {
			It("should return method allowed", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/student", nil)
				handler := main.RequestMethodGet(main.StudentHandler())
				handler.ServeHTTP(w, r)
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(w.Body.String()).To(Equal("Welcome to Student page"))
			})
		})

		When("hit with POST request method", func() {
			It("should return method not allowed", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("POST", "/student", nil)
				handler := main.RequestMethodGet(main.StudentHandler())
				handler.ServeHTTP(w, r)
				Expect(w.Code).To(Equal(http.StatusMethodNotAllowed))
				Expect(w.Body.String()).To(Equal("Method is not allowed"))
			})
		})

		When("hit with PUT request method", func() {
			It("should return method not allowed", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("PUT", "/student", nil)
				handler := main.RequestMethodGet(main.StudentHandler())
				handler.ServeHTTP(w, r)
				Expect(w.Code).To(Equal(http.StatusMethodNotAllowed))
				Expect(w.Body.String()).To(Equal("Method is not allowed"))
			})
		})

		When("hit with DELETE request method", func() {
			It("should return method not allowed", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("DELETE", "/student", nil)
				handler := main.RequestMethodGet(main.StudentHandler())
				handler.ServeHTTP(w, r)
				Expect(w.Code).To(Equal(http.StatusMethodNotAllowed))
				Expect(w.Body.String()).To(Equal("Method is not allowed"))
			})
		})
	})
})
