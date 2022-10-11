package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/model"
)

var _ = Describe("Api Todo List", func() {
	var server main.API
	BeforeEach(func() {
		server = main.NewAPI()
	})

	Describe("/user/register", func() {
		When("hit endpoint with GET method", func() {
			It("should return a wrong method", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodGet, "/user/register", nil)
				server.Handler().ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusMethodNotAllowed))
				Expect(ErrorResponse.Error).To(Equal("Method is not allowed!"))
			})
		})

		When("send empty request body with POST method", func() {
			It("should return a wrong body request", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/user/register", nil)
				server.Handler().ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusBadRequest))
				Expect(ErrorResponse.Error).To(Equal("Internal Server Error"))
			})
		})

		When("send empty username and password with POST method", func() {
			It("should return a wrong method", func() {
				cred := model.Credentials{Username: "", Password: ""}
				register, _ := json.Marshal(cred)
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/user/register", bytes.NewReader(register))
				server.Handler().ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusBadRequest))
				Expect(ErrorResponse.Error).To(Equal("Username or Password empty"))
			})
		})

		When("send username and password with POST method", func() {
			It("should return a successful register response", func() {
				cred := model.Credentials{Username: "aditira", Password: "12345"}
				register, _ := json.Marshal(cred)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/user/register", bytes.NewReader(register))
				server.Handler().ServeHTTP(w, r)
				SuccessResponse := model.SuccessResponse{}
				json.NewDecoder(w.Body).Decode(&SuccessResponse)

				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
				Expect(SuccessResponse.Username).To(Equal("aditira"))
				Expect(SuccessResponse.Message).To(Equal("Register Success"))

				if value, ok := db.Users[SuccessResponse.Username]; ok {
					Expect(ok).To(BeTrue())
					Expect(value).To(Equal("12345"))
				}
			})
		})

		When("send the same username twice by POST method", func() {
			It("should return a conflict response", func() {
				cred := model.Credentials{Username: "aditira", Password: "12345"}
				register, _ := json.Marshal(cred)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/user/register", bytes.NewReader(register))
				server.Handler().ServeHTTP(w, r)

				w = httptest.NewRecorder()
				r = httptest.NewRequest(http.MethodPost, "/user/register", bytes.NewReader(register))
				server.Handler().ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusConflict))
				Expect(ErrorResponse.Error).To(Equal("Username already exist"))

			})
		})
	})

	Describe("/user/login", func() {
		When("hit endpoint with GET method", func() {
			It("should return a wrong method", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodGet, "/user/login", nil)
				server.Handler().ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusMethodNotAllowed))
				Expect(ErrorResponse.Error).To(Equal("Method is not allowed!"))
			})
		})

		When("send empty request body with POST method", func() {
			It("should return a wrong body request", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/user/login", nil)
				server.Handler().ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusBadRequest))
				Expect(ErrorResponse.Error).To(Equal("Internal Server Error"))
			})
		})

		When("send empty username and password with POST method", func() {
			It("should return a wrong method", func() {
				cred := model.Credentials{Username: "", Password: ""}
				register, _ := json.Marshal(cred)
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/user/login", bytes.NewReader(register))
				server.Handler().ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusBadRequest))
				Expect(ErrorResponse.Error).To(Equal("Username or Password empty"))
			})
		})

		When("the username and password are incorrect with POST method", func() {
			It("should return unauthorized", func() {
				cred := model.Credentials{Username: "dito", Password: "12345"}
				register, _ := json.Marshal(cred)

				//register
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/user/register", bytes.NewReader(register))
				server.Handler().ServeHTTP(w, r)

				cred = model.Credentials{Username: "adit", Password: "123"}
				login, _ := json.Marshal(cred)

				//login
				w = httptest.NewRecorder()
				r = httptest.NewRequest(http.MethodPost, "/user/login", bytes.NewReader(login))
				server.Handler().ServeHTTP(w, r)

				cookies := w.Result().Cookies()
				var isCookieTokenExist bool
				for _, c := range cookies {
					if c.Name == "session_token" {
						isCookieTokenExist = true
						break
					}
				}
				Expect(isCookieTokenExist).To(BeFalse())
				Expect(w.Result().StatusCode).To(Equal(http.StatusUnauthorized))

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(ErrorResponse.Error).To(Equal("Wrong User or Password!"))
			})
		})

		When("send username and password are correct with POST method", func() {
			It("should return a successful login response", func() {
				cred := model.Credentials{Username: "aditira", Password: "12345"}
				login, _ := json.Marshal(cred)

				//register
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/user/register", bytes.NewReader(login))
				server.Handler().ServeHTTP(w, r)

				//login
				w = httptest.NewRecorder()
				r = httptest.NewRequest(http.MethodPost, "/user/login", bytes.NewReader(login))
				server.Handler().ServeHTTP(w, r)
				SuccessResponse := model.SuccessResponse{}
				json.NewDecoder(w.Body).Decode(&SuccessResponse)

				cookies := w.Result().Cookies()
				var isCookieTokenExist bool
				for _, c := range cookies {
					if c.Name == "session_token" {
						isCookieTokenExist = true
						break
					}
				}
				Expect(isCookieTokenExist).To(BeTrue())

				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
				Expect(SuccessResponse.Username).To(Equal("aditira"))
				Expect(SuccessResponse.Message).To(Equal("Login Success"))
			})
		})
	})

	Describe("/todo/clear", func() {
		When("clear todo list with GET method", func() {
			It("should return a wrong method", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodGet, "/todo/clear", nil)
				server.Handler().ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusMethodNotAllowed))
				Expect(ErrorResponse.Error).To(Equal("Method is not allowed!"))
			})
		})

		When("clear todo list with DELETE method, without login cookie", func() {
			It("should return a cookie not found", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodDelete, "/todo/clear", nil)
				server.Handler().ServeHTTP(w, r)

				err := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&err)

				Expect(w.Result().StatusCode).To(Equal(http.StatusUnauthorized))
				Expect(err.Error).To(Equal("http: named cookie not present"))
			})
		})

		When("clear todo list with DELETE method and login cookie", func() {
			It("should return a successful clear todo list response", func() {
				cred := model.Credentials{Username: "aditira", Password: "12345"}
				login, _ := json.Marshal(cred)

				//register
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/user/register", bytes.NewReader(login))
				server.Handler().ServeHTTP(w, r)

				//login
				w = httptest.NewRecorder()
				r = httptest.NewRequest(http.MethodPost, "/user/login", bytes.NewReader(login))
				server.Handler().ServeHTTP(w, r)

				var cookie *http.Cookie
				for _, c := range w.Result().Cookies() {
					if c.Name == "session_token" {
						cookie = c
					}
				}

				//clear
				w = httptest.NewRecorder()
				r = httptest.NewRequest(http.MethodDelete, "/todo/clear", nil)
				r.AddCookie(cookie)
				server.Handler().ServeHTTP(w, r)

				todoRes := model.SuccessResponse{}
				json.NewDecoder(w.Body).Decode(&todoRes)

				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
				Expect(todoRes.Username).To(Equal("aditira"))
				Expect(todoRes.Message).To(Equal("Clear ToDo Success"))
			})
		})
	})

	Describe("/todo/create", func() {
		When("create todo list with GET method", func() {
			It("should return a wrong method", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodGet, "/todo/create", nil)
				server.Handler().ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusMethodNotAllowed))
				Expect(ErrorResponse.Error).To(Equal("Method is not allowed!"))
			})
		})

		When("create todo list with POST method, without login cookie", func() {
			It("should return a cookie not found", func() {
				//create todo
				todo := model.Todo{Task: "Create a todo app program", Done: false}
				todoBody, _ := json.Marshal(todo)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/todo/create", bytes.NewReader(todoBody))
				server.Handler().ServeHTTP(w, r)

				err := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&err)

				Expect(w.Result().StatusCode).To(Equal(http.StatusUnauthorized))
				Expect(err.Error).To(Equal("http: named cookie not present"))
			})
		})

		When("create todo list with POST method and login cookie", func() {
			It("should return a successful create todo list response", func() {
				cred := model.Credentials{Username: "aditira", Password: "12345"}
				login, _ := json.Marshal(cred)

				//register
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/user/register", bytes.NewReader(login))
				server.Handler().ServeHTTP(w, r)

				//login
				w = httptest.NewRecorder()
				r = httptest.NewRequest(http.MethodPost, "/user/login", bytes.NewReader(login))
				server.Handler().ServeHTTP(w, r)

				var cookie *http.Cookie
				for _, c := range w.Result().Cookies() {
					if c.Name == "session_token" {
						cookie = c
					}
				}

				//create todo
				todo := model.Todo{Task: "Create a todo app program", Done: false}
				todoBody, _ := json.Marshal(todo)

				w = httptest.NewRecorder()
				r = httptest.NewRequest(http.MethodPost, "/todo/create", bytes.NewReader(todoBody))
				r.AddCookie(cookie)
				server.Handler().ServeHTTP(w, r)

				todoRes := model.SuccessResponse{}
				json.NewDecoder(w.Body).Decode(&todoRes)

				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
				Expect(todoRes.Username).To(Equal("aditira"))
				Expect(todoRes.Message).To(Equal(fmt.Sprintf("Task %s added!", todo.Task)))
			})
		})
	})

	Describe("/todo/read", func() {
		When("read todo list with POST method", func() {
			It("should return a wrong method", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/todo/read", nil)
				server.Handler().ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusMethodNotAllowed))
				Expect(ErrorResponse.Error).To(Equal("Method is not allowed!"))
			})
		})

		When("read todo list with GET method, without login cookie", func() {
			It("should return a cookie not found", func() {
				//read todo
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodGet, "/todo/read", nil)
				server.Handler().ServeHTTP(w, r)

				err := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&err)

				Expect(w.Result().StatusCode).To(Equal(http.StatusUnauthorized))
				Expect(err.Error).To(Equal("http: named cookie not present"))
			})
		})

		When("read task list with GET method and login cookie but the task database is empty", func() {
			It("should return not found", func() {
				cred := model.Credentials{Username: "aditira", Password: "12345"}
				login, _ := json.Marshal(cred)

				//register
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/user/register", bytes.NewReader(login))
				server.Handler().ServeHTTP(w, r)

				//login
				w = httptest.NewRecorder()
				r = httptest.NewRequest(http.MethodPost, "/user/login", bytes.NewReader(login))
				server.Handler().ServeHTTP(w, r)

				var cookie *http.Cookie
				for _, c := range w.Result().Cookies() {
					if c.Name == "session_token" {
						cookie = c
					}
				}

				//reset
				w = httptest.NewRecorder()
				r = httptest.NewRequest(http.MethodDelete, "/todo/reset", nil)
				r.AddCookie(cookie)
				server.Handler().ServeHTTP(w, r)

				//read todo
				w = httptest.NewRecorder()
				r = httptest.NewRequest(http.MethodGet, "/todo/read", nil)
				r.AddCookie(cookie)
				server.Handler().ServeHTTP(w, r)

				err := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&err)

				Expect(w.Result().StatusCode).To(Equal(http.StatusNotFound))
				Expect(err.Error).To(Equal("Todolist not found!"))
			})
		})

		When("read task list with GET method and login cookie but no task found from logged in user", func() {
			It("should return empty todo list response", func() {
				cred := model.Credentials{Username: "aditira", Password: "12345"}
				login, _ := json.Marshal(cred)

				//register
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/user/register", bytes.NewReader(login))
				server.Handler().ServeHTTP(w, r)

				//login
				w = httptest.NewRecorder()
				r = httptest.NewRequest(http.MethodPost, "/user/login", bytes.NewReader(login))
				server.Handler().ServeHTTP(w, r)

				var cookie *http.Cookie
				for _, c := range w.Result().Cookies() {
					if c.Name == "session_token" {
						cookie = c
					}
				}

				//clear
				w = httptest.NewRecorder()
				r = httptest.NewRequest(http.MethodDelete, "/todo/clear", nil)
				r.AddCookie(cookie)
				server.Handler().ServeHTTP(w, r)

				//read todo
				w = httptest.NewRecorder()
				r = httptest.NewRequest(http.MethodGet, "/todo/read", nil)
				r.AddCookie(cookie)
				server.Handler().ServeHTTP(w, r)

				todoRes := []model.Todo{}
				json.NewDecoder(w.Body).Decode(&todoRes)

				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
				Expect(len(todoRes)).To(Equal(0))
			})
		})

		When("read todo list with GET method, login cookie and create todo list", func() {
			It("should return created todo list", func() {
				cred := model.Credentials{Username: "aditira", Password: "12345"}
				login, _ := json.Marshal(cred)

				//register
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/user/register", bytes.NewReader(login))
				server.Handler().ServeHTTP(w, r)

				//login
				w = httptest.NewRecorder()
				r = httptest.NewRequest(http.MethodPost, "/user/login", bytes.NewReader(login))
				server.Handler().ServeHTTP(w, r)

				var cookie *http.Cookie
				for _, c := range w.Result().Cookies() {
					if c.Name == "session_token" {
						cookie = c
					}
				}

				//clear
				w = httptest.NewRecorder()
				r = httptest.NewRequest(http.MethodDelete, "/todo/clear", nil)
				r.AddCookie(cookie)
				server.Handler().ServeHTTP(w, r)

				//create todo
				todo := model.Todo{Task: "Create a todo app program", Done: false}
				todoBody, _ := json.Marshal(todo)

				w = httptest.NewRecorder()
				r = httptest.NewRequest(http.MethodPost, "/todo/create", bytes.NewReader(todoBody))
				r.AddCookie(cookie)
				server.Handler().ServeHTTP(w, r)

				//read todo
				w = httptest.NewRecorder()
				r = httptest.NewRequest(http.MethodGet, "/todo/read", nil)
				r.AddCookie(cookie)
				server.Handler().ServeHTTP(w, r)

				todoRes := []model.Todo{}
				json.NewDecoder(w.Body).Decode(&todoRes)

				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
				Expect(len(todoRes)).To(Equal(1))
				Expect(todoRes[0].Task).To(Equal("Create a todo app program"))
				Expect(todoRes[0].Done).To(Equal(false))
			})
		})
	})

	Describe("/user/logout", func() {
		When("logout user with POST method", func() {
			It("should return a wrong method", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/user/logout", nil)
				server.Handler().ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusMethodNotAllowed))
				Expect(ErrorResponse.Error).To(Equal("Method is not allowed!"))
			})
		})

		When("logout user with GET method, without login cookie", func() {
			It("should return a cookie not found", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodGet, "/user/logout", nil)
				server.Handler().ServeHTTP(w, r)

				err := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&err)

				Expect(w.Result().StatusCode).To(Equal(http.StatusUnauthorized))
				Expect(err.Error).To(Equal("http: named cookie not present"))
			})
		})

		When("logout user with GET method and login cookie", func() {
			It("should return a logout success", func() {
				cred := model.Credentials{Username: "aditira", Password: "12345"}
				login, _ := json.Marshal(cred)

				//register
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/user/register", bytes.NewReader(login))
				server.Handler().ServeHTTP(w, r)

				//login
				w = httptest.NewRecorder()
				r = httptest.NewRequest(http.MethodPost, "/user/login", bytes.NewReader(login))
				server.Handler().ServeHTTP(w, r)

				var cookie *http.Cookie
				for _, c := range w.Result().Cookies() {
					if c.Name == "session_token" {
						cookie = c
					}
				}

				//logout
				w = httptest.NewRecorder()
				r = httptest.NewRequest(http.MethodGet, "/user/logout", nil)
				r.AddCookie(cookie)
				server.Handler().ServeHTTP(w, r)

				todoRes := model.SuccessResponse{}
				json.NewDecoder(w.Body).Decode(&todoRes)

				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
				Expect(todoRes.Username).To(Equal("aditira"))
				Expect(todoRes.Message).To(Equal("Logout Success"))
			})
		})
	})
})
