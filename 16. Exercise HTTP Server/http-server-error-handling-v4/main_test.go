package main_test

import (
	main "a21hc3NpZ25tZW50"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Error Handling", func() {
	Describe("hit endpoint / with POST request method", func() {
		It("should return status code 405 and message `Method not allowed`", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/", nil)
			main.MethodHandler().ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusMethodNotAllowed))
			Expect(w.Body.String()).To(Equal("Method not allowed"))
		})
	})

	Describe("hit endpoint / with GET request method", func() {
		It("should return status code 200 and message `Method handler passed`", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/", nil)
			main.MethodHandler().ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("Method handler passed"))
		})
	})

	Describe("hit endpoint / with empty parameter data", func() {
		It("should return status code 404 and message `Data not found`", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/", nil)
			main.DataHandler().ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusNotFound))
			Expect(w.Body.String()).To(Equal("Data not found"))
		})
	})

	Describe("hit endpoint / with parameter data filled", func() {
		It("should return status code 200 and message `Data handler passed`", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/?data=Aditira", nil)
			main.DataHandler().ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("Data handler passed"))
		})
	})

	Describe("hit endpoint / with empty parameter filename", func() {
		It("should return status code 500 and message `File not found`", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/", nil)
			main.OpenFileHandler().ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))
			Expect(w.Body.String()).To(Equal("File not found"))
		})
	})

	Describe("hit endpoint / with parameter filename is wrong.txt", func() {
		It("should return status code 500 and message `File not found`", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/?filename=wrong.txt", nil)
			main.OpenFileHandler().ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))
			Expect(w.Body.String()).To(Equal("File not found"))
		})
	})

	Describe("hit endpoint / with parameter filename is hello.txt", func() {
		It("should return status code 200 and message `Error handler passed`", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/?filename=hello.txt", nil)
			main.OpenFileHandler().ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("Error handler passed"))
		})
	})
})
