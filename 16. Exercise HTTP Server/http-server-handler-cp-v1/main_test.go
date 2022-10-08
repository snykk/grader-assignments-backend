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

var _ = Describe("Basic Handler", func() {
	Describe("hit endpoint / with GET request method", func() {
		It("should return current day and time", func() {
			t := time.Now()
			expected := fmt.Sprintf("%v, %v %v %v", t.Weekday(), t.Day(), t.Month(), t.Year())

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/", nil)
			main.GetHandler().ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal(expected))
		})
	})
})
