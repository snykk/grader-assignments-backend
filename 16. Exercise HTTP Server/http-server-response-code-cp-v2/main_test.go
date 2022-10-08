package main_test

import (
	main "a21hc3NpZ25tZW50"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Http Server Reponse Code", func() {
	Describe("hit endpoint /students with GET request method and non exists name", func() {
		It("should return status not found", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/students?name=Sophia", nil)
			main.GetMux().ServeHTTP(w, r)
			body, err := io.ReadAll(w.Body)
			Expect(err).To(BeNil())
			Expect(w.Code).To(Equal(http.StatusNotFound))
			Expect(string(body)).To(Equal("Data not found"))
		})
	})

	Describe("hit endpoint /students with GET request method and exists name", func() {
		When("name is Aditira", func() {
			It("should return status OK", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/students?name=Aditira", nil)
				fmt.Println(r.URL.Query())
				main.GetMux().ServeHTTP(w, r)
				body, err := io.ReadAll(w.Body)
				Expect(err).To(BeNil())
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(string(body)).To(Equal("Name is exists"))
			})
		})
		When("name is Dito", func() {
			It("should return status OK", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/students?name=Dito", nil)
				fmt.Println(r.URL.Query())
				main.GetMux().ServeHTTP(w, r)
				body, err := io.ReadAll(w.Body)
				Expect(err).To(BeNil())
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(string(body)).To(Equal("Name is exists"))
			})
		})
		When("name is Afis", func() {
			It("should return status OK", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/students?name=Afis", nil)
				fmt.Println(r.URL.Query())
				main.GetMux().ServeHTTP(w, r)
				body, err := io.ReadAll(w.Body)
				Expect(err).To(BeNil())
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(string(body)).To(Equal("Name is exists"))
			})
		})
		When("name is Eddy", func() {
			It("should return status OK", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/students?name=Eddy", nil)
				fmt.Println(r.URL.Query())
				main.GetMux().ServeHTTP(w, r)
				body, err := io.ReadAll(w.Body)
				Expect(err).To(BeNil())
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(string(body)).To(Equal("Name is exists"))
			})
		})
	})

	Describe("hit endpoint /students with POST request method", func() {
		It("should return status method not allowed", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/students?name=Aditira", nil)
			main.GetMux().ServeHTTP(w, r)
			body, err := io.ReadAll(w.Body)
			Expect(err).To(BeNil())
			Expect(w.Code).To(Equal(http.StatusMethodNotAllowed))
			Expect(string(body)).To(Equal("Method is not allowed"))
		})
	})
})
