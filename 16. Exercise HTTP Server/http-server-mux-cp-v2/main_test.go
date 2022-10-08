package main_test

import (
	main "a21hc3NpZ25tZW50"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handler", func() {
	Describe("hit endpoint /time with GET request method", func() {
		It("should return current day and time", func() {
			t := time.Now()
			expected := fmt.Sprintf("%v, %v %v %v", t.Weekday(), t.Day(), t.Month(), t.Year())

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/time", nil)
			main.GetMux().ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal(expected))
		})
	})

	Describe("hit endpoint /hello with GET request method", func() {
		It("should return Hello there", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/hello", nil)
			main.GetMux().ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("Hello there"))
		})
	})
	Describe("hit endpoint /hello with GET request method and query param", func() {
		When("query param is Aditira", func() {
			It("should return Hello, Aditira!", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/hello?name=Aditira", nil)
				main.GetMux().ServeHTTP(w, r)
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(w.Body.String()).To(Equal("Hello, Aditira!"))
			})
		})
		When("query param is Dito", func() {
			It("should return Hello, Dito!", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/hello?name=Dito", nil)
				main.GetMux().ServeHTTP(w, r)
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(w.Body.String()).To(Equal("Hello, Dito!"))
			})
		})
		When("query param is Afis", func() {
			It("should return Hello, Afis!", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/hello?name=Afis", nil)
				main.GetMux().ServeHTTP(w, r)
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(w.Body.String()).To(Equal("Hello, Afis!"))
			})
		})
		When("query param is Eddy", func() {
			It("should return Hello, Eddy!", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/hello?name=Eddy", nil)
				main.GetMux().ServeHTTP(w, r)
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(w.Body.String()).To(Equal("Hello, Eddy!"))
			})
		})
	})
})
