package main_test

import (
	main "a21hc3NpZ25tZW50"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Custom Handler", func() {
	Describe("hit endpoint / with GET request method", func() {
		It("should return current day and time", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/", nil)
			handler := main.QuotesHandler{}
			handler.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(main.Quotes).Should(ContainElement(w.Body.String()))
		})
	})
	Describe("hit endpoint / with POST request method", func() {
		It("should return current day and time", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/", nil)
			handler := main.QuotesHandler{}
			handler.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(main.Quotes).Should(ContainElement(w.Body.String()))
		})
	})
	Describe("hit endpoint /quote with GET request method", func() {
		It("should return current day and time", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/quote", nil)
			handler := main.QuotesHandler{}
			handler.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(main.Quotes).Should(ContainElement(w.Body.String()))
		})
	})
})
