package main_test

import (
	"a21hc3NpZ25tZW50/api"
	"a21hc3NpZ25tZW50/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"a21hc3NpZ25tZW50/db"
	repo "a21hc3NpZ25tZW50/repository"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cashier App", func() {
	var server api.API
	var userRepo repo.UserRepository
	var sessionRepo repo.SessionsRepository
	var productRepo repo.ProductRepository
	var cartRepo repo.CartRepository
	BeforeEach(func() {
		db := &db.JsonDB{}
		userRepo = repo.NewUserRepository(db)
		sessionRepo = repo.NewSessionsRepository(db)
		productRepo = repo.NewProductRepository(db)
		cartRepo = repo.NewCartRepository(db)

		server = api.NewAPI(userRepo, sessionRepo, productRepo, cartRepo)

		userRepo.ResetUser()
		sessionRepo.ResetSessions()
		cartRepo.ResetCarts()
	})

	Describe("Repository", func() {
		Describe("Users repository", func() {
			When("add user data to file users.json", func() {
				It("should save data user as a list to file users.json", func() {
					err := userRepo.AddUser(model.Credentials{
						Username: "aditira",
						Password: "12345",
					})
					Expect(err).ShouldNot(HaveOccurred())

					listUsers, err := userRepo.ReadUser()
					expectListUsers := []model.Credentials{
						{Username: "aditira", Password: "12345"},
					}
					Expect(err).ShouldNot(HaveOccurred())
					Expect(listUsers).To(Equal(expectListUsers))

					userRepo.ResetUser()
					sessionRepo.ResetSessions()
					cartRepo.ResetCarts()
				})
			})
		})

		Describe("Sessions repository", func() {
			When("add session data to file sessions.json", func() {
				It("should save data session as a list to file sessions.json", func() {
					session := model.Session{
						Token:    "cc03dbea-4085-47ba-86fe-020f5d01a9d8",
						Username: "aditira",
						Expiry:   time.Date(2022, 11, 17, 20, 34, 58, 651387237, time.UTC),
					}
					err := sessionRepo.AddSessions(session)
					Expect(err).ShouldNot(HaveOccurred())

					listSessions, err := sessionRepo.ReadSessions()
					Expect(err).ShouldNot(HaveOccurred())
					token, err := sessionRepo.TokenExist(listSessions, "cc03dbea-4085-47ba-86fe-020f5d01a9d8")
					Expect(err).ShouldNot(HaveOccurred())
					Expect(token).To(Equal(session))
					Expect(listSessions[0]).To(Equal(session))

					sessionRepo.ResetSessions()
				})
			})

			When("delete selected session to file sessions.json", func() {
				It("should delete data session target from list in file sessions.json", func() {
					session := model.Session{
						Token:    "cc03dbea-4085-47ba-86fe-020f5d01a9d8",
						Username: "aditira",
						Expiry:   time.Date(2022, 11, 17, 20, 34, 58, 651387237, time.UTC),
					}
					err := sessionRepo.AddSessions(session)
					Expect(err).ShouldNot(HaveOccurred())

					listSessions, err := sessionRepo.ReadSessions()
					Expect(err).ShouldNot(HaveOccurred())

					token, err := sessionRepo.TokenExist(listSessions, "cc03dbea-4085-47ba-86fe-020f5d01a9d8")
					Expect(err).ShouldNot(HaveOccurred())
					Expect(token).To(Equal(session))

					err = sessionRepo.DeleteSessions("cc03dbea-4085-47ba-86fe-020f5d01a9d8")
					Expect(err).ShouldNot(HaveOccurred())

					listSessions, err = sessionRepo.ReadSessions()
					Expect(err).ShouldNot(HaveOccurred())

					token, err = sessionRepo.TokenExist(listSessions, "cc03dbea-4085-47ba-86fe-020f5d01a9d8")
					Expect(err).To(Equal(fmt.Errorf("Token Not Found!")))
					Expect(token).To(Equal(model.Session{}))

					sessionRepo.ResetSessions()
				})
			})

			When("check expired session with exprired session is 5 hours from now", func() {
				It("should return a session model with token, name, and expired time", func() {
					session := model.Session{
						Token:    "cc03dbea-4085-47ba-86fe-020f5d01a9d8",
						Username: "aditira",
						Expiry:   time.Now().Add(5 * time.Hour),
					}
					err := sessionRepo.AddSessions(session)
					Expect(err).ShouldNot(HaveOccurred())

					tokenFound, err := sessionRepo.CheckExpireToken("cc03dbea-4085-47ba-86fe-020f5d01a9d8")
					Expect(err).ShouldNot(HaveOccurred())
					Expect(tokenFound.Token).To(Equal("cc03dbea-4085-47ba-86fe-020f5d01a9d8"))
					Expect(tokenFound.Username).To(Equal("aditira"))

					sessionRepo.ResetSessions()
				})
			})

			When("check expired session with exprired session is 5 hours ago", func() {
				It("should return error message token is expired and empty session model", func() {
					session := model.Session{
						Token:    "cc03dbea-4085-47ba-86fe-020f5d01a9d8",
						Username: "aditira",
						Expiry:   time.Now().Add(-5 * time.Hour),
					}
					err := sessionRepo.AddSessions(session)
					Expect(err).ShouldNot(HaveOccurred())

					tokenFound, err := sessionRepo.CheckExpireToken("cc03dbea-4085-47ba-86fe-020f5d01a9d8")
					Expect(err).To(Equal(fmt.Errorf("Token is Expired!")))
					Expect(tokenFound).To(Equal(model.Session{}))

					sessionRepo.ResetSessions()
				})
			})
		})

		Describe("Product repository", func() {
			When("read list data product", func() {
				It("should return list product", func() {
					listProduct, err := productRepo.ReadProducts()
					Expect(err).ShouldNot(HaveOccurred())

					expectList := []model.Product{
						{Id: "1", Name: "Tea", Price: "20000", Quantity: "", Total: 0},
						{Id: "2", Name: "Milk", Price: "10000", Quantity: "", Total: 0},
						{Id: "3", Name: "Cofe", Price: "25000", Quantity: "", Total: 0},
						{Id: "4", Name: "Apple", Price: "15000", Quantity: "", Total: 0},
						{Id: "5", Name: "Bread", Price: "30000", Quantity: "", Total: 0},
						{Id: "6", Name: "Cake", Price: "40000", Quantity: "", Total: 0},
					}
					Expect(listProduct).To(Equal(expectList))
				})
			})
		})

		Describe("Cart repository", func() {
			When("add cart data to file carts.json", func() {
				It("should save data cart to file carts.json", func() {
					productCart := model.Cart{
						Name: "aditira",
						Cart: []model.Product{
							{Id: "1", Name: "Tea", Price: "20000", Quantity: "1", Total: 20000},
							{Id: "2", Name: "Milk", Price: "10000", Quantity: "2", Total: 20000},
							{Id: "3", Name: "Cofe", Price: "25000", Quantity: "3", Total: 75000},
						},
						TotalPrice: 115000,
					}
					err := cartRepo.AddCart(productCart)
					Expect(err).ShouldNot(HaveOccurred())

					expectProdCart, err := cartRepo.ReadCart()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(productCart).To(Equal(expectProdCart))

					cartRepo.ResetCarts()
				})
			})
		})

	})

	Describe("API", func() {
		Describe("/page/register", func() {
			When("call endpoint with POST method", func() {
				It("should return a wrong method", func() {
					w := httptest.NewRecorder()
					r := httptest.NewRequest(http.MethodPost, "/page/register", nil)
					server.Handler().ServeHTTP(w, r)

					ErrorResponse := model.ErrorResponse{}
					json.NewDecoder(w.Body).Decode(&ErrorResponse)
					Expect(w.Result().StatusCode).To(Equal(http.StatusMethodNotAllowed))
					Expect(ErrorResponse.Error).To(Equal("Method is not allowed!"))
				})
			})
			When("call endpoint with GET method", func() {
				It("should return file register.html on views folder as response", func() {
					w := httptest.NewRecorder()
					r := httptest.NewRequest(http.MethodGet, "/page/register", nil)
					server.Handler().ServeHTTP(w, r)
					register, err := ioutil.ReadFile("./views/register.html")
					Expect(err).ShouldNot(HaveOccurred())
					Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
					Expect(w.Body.String()).To(Equal(string(register)))
				})
			})
		})

		Describe("/page/login", func() {
			When("call endpoint with POST method", func() {
				It("should return a wrong method", func() {
					w := httptest.NewRecorder()
					r := httptest.NewRequest(http.MethodPost, "/page/login", nil)
					server.Handler().ServeHTTP(w, r)

					ErrorResponse := model.ErrorResponse{}
					json.NewDecoder(w.Body).Decode(&ErrorResponse)
					Expect(w.Result().StatusCode).To(Equal(http.StatusMethodNotAllowed))
					Expect(ErrorResponse.Error).To(Equal("Method is not allowed!"))
				})
			})

			When("call endpoint with GET method", func() {
				It("should return file login.html on views folder as response", func() {
					w := httptest.NewRecorder()
					r := httptest.NewRequest(http.MethodGet, "/page/login", nil)
					server.Handler().ServeHTTP(w, r)
					login, err := ioutil.ReadFile("./views/login.html")
					Expect(err).ShouldNot(HaveOccurred())
					Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
					Expect(w.Body.String()).To(Equal(string(login)))
				})
			})
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
					Expect(ErrorResponse.Error).To(Equal("Username or Password empty"))
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
					w := httptest.NewRecorder()
					r := httptest.NewRequest(http.MethodPost, "/user/register", strings.NewReader("username=aditira&password=12345"))
					r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					server.Handler().ServeHTTP(w, r)
					Expect(w.Result().StatusCode).To(Equal(http.StatusOK))

					listUser, err := userRepo.ReadUser()
					Expect(err).ShouldNot(HaveOccurred())
					creds := model.Credentials{
						Username: "aditira",
						Password: "12345",
					}
					Expect(userRepo.LoginValid(listUser, creds)).To(BeTrue())

					userRepo.ResetUser()
					sessionRepo.ResetSessions()
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
					Expect(ErrorResponse.Error).To(Equal("Username or Password empty"))
				})
			})

			When("send empty username and password with POST method", func() {
				It("should return a username or password empty", func() {
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

			When("send username and password are correct with POST method", func() {
				It("should return a successful login response", func() {
					//register
					w := httptest.NewRecorder()
					r := httptest.NewRequest(http.MethodPost, "/user/register", strings.NewReader("username=aditira&password=12345"))
					r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					server.Handler().ServeHTTP(w, r)

					listUser, err := userRepo.ReadUser()
					Expect(err).ShouldNot(HaveOccurred())
					creds := model.Credentials{
						Username: "aditira",
						Password: "12345",
					}
					Expect(userRepo.LoginValid(listUser, creds)).To(BeTrue())

					//login
					w = httptest.NewRecorder()
					r = httptest.NewRequest(http.MethodPost, "/user/login", strings.NewReader("username=aditira&password=12345"))
					r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					server.Handler().ServeHTTP(w, r)

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

					userRepo.ResetUser()
					sessionRepo.ResetSessions()
				})
			})

			When("the username and password are incorrect with POST method", func() {
				It("should return unauthorized", func() {
					//register
					w := httptest.NewRecorder()
					r := httptest.NewRequest(http.MethodPost, "/user/register", strings.NewReader("username=dito&password=12345"))
					r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					server.Handler().ServeHTTP(w, r)

					listUser, err := userRepo.ReadUser()
					Expect(err).ShouldNot(HaveOccurred())
					creds := model.Credentials{
						Username: "dito",
						Password: "12345",
					}
					Expect(userRepo.LoginValid(listUser, creds)).To(BeTrue())

					//login
					w = httptest.NewRecorder()
					r = httptest.NewRequest(http.MethodPost, "/user/login", strings.NewReader("username=adit&password=123"))
					r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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

					ErrorResponse := model.ErrorResponse{}
					json.NewDecoder(w.Body).Decode(&ErrorResponse)
					Expect(w.Result().StatusCode).To(Equal(http.StatusUnauthorized))
					Expect(ErrorResponse.Error).To(Equal("Wrong User or Password!"))

					userRepo.ResetUser()
					sessionRepo.ResetSessions()
				})
			})
		})

		Describe("/cart/add", func() {
			When("add product to cart with GET method", func() {
				It("should return a wrong method", func() {
					w := httptest.NewRecorder()
					r := httptest.NewRequest(http.MethodGet, "/cart/add", nil)
					server.Handler().ServeHTTP(w, r)

					ErrorResponse := model.ErrorResponse{}
					json.NewDecoder(w.Body).Decode(&ErrorResponse)
					Expect(w.Result().StatusCode).To(Equal(http.StatusMethodNotAllowed))
					Expect(ErrorResponse.Error).To(Equal("Method is not allowed!"))
				})
			})

			When("add product to cart with POST method, but without login cookie", func() {
				It("should return a wrong method", func() {
					w := httptest.NewRecorder()
					r := httptest.NewRequest(http.MethodPost, "/cart/add", nil)
					server.Handler().ServeHTTP(w, r)

					err := model.ErrorResponse{}
					json.NewDecoder(w.Body).Decode(&err)

					Expect(w.Result().StatusCode).To(Equal(http.StatusUnauthorized))
					Expect(err.Error).To(Equal("http: named cookie not present"))
				})
			})

			When("send empty request body with POST method and login cookie", func() {
				It("should return a product request is not found", func() {
					//register
					w := httptest.NewRecorder()
					r := httptest.NewRequest(http.MethodPost, "/user/register", strings.NewReader("username=aditira&password=12345"))
					r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					server.Handler().ServeHTTP(w, r)
					Expect(w.Result().StatusCode).To(Equal(http.StatusOK))

					//login
					w = httptest.NewRecorder()
					r = httptest.NewRequest(http.MethodPost, "/user/login", strings.NewReader("username=aditira&password=12345"))
					r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					server.Handler().ServeHTTP(w, r)

					var cookie *http.Cookie
					var isCookieTokenExist bool
					for _, c := range w.Result().Cookies() {
						if c.Name == "session_token" {
							isCookieTokenExist = true
							cookie = c
						}
					}
					Expect(isCookieTokenExist).To(BeTrue())

					//add product without request body.
					w = httptest.NewRecorder()
					r = httptest.NewRequest(http.MethodPost, "/cart/add", nil)
					r.AddCookie(cookie)
					server.Handler().ServeHTTP(w, r)

					err := model.ErrorResponse{}
					json.NewDecoder(w.Body).Decode(&err)

					Expect(w.Result().StatusCode).To(Equal(http.StatusBadRequest))
					Expect(err.Error).To(Equal("Request Product Not Found"))

					userRepo.ResetUser()
					sessionRepo.ResetSessions()
					cartRepo.ResetCarts()
				})
			})

			When("add 1 product to cart with POST method and login cookie", func() {
				It("should added 1 cart data to file carts.json with `quantity per-product`, `price per-product` and `total price` calculated", func() {
					//register
					w := httptest.NewRecorder()
					r := httptest.NewRequest(http.MethodPost, "/user/register", strings.NewReader("username=aditira&password=12345"))
					r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					server.Handler().ServeHTTP(w, r)

					//login
					w = httptest.NewRecorder()
					r = httptest.NewRequest(http.MethodPost, "/user/login", strings.NewReader("username=aditira&password=12345"))
					r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					server.Handler().ServeHTTP(w, r)

					var cookie *http.Cookie
					var isCookieTokenExist bool
					for _, c := range w.Result().Cookies() {
						if c.Name == "session_token" {
							isCookieTokenExist = true
							cookie = c
						}
					}
					Expect(isCookieTokenExist).To(BeTrue())

					//add 1 product to cart
					w = httptest.NewRecorder()
					r = httptest.NewRequest(http.MethodPost, "/cart/add", strings.NewReader("product=1,Tea,20000,3"))
					r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					r.AddCookie(cookie)
					server.Handler().ServeHTTP(w, r)

					Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
					cart, err := cartRepo.ReadCart()
					Expect(err).ShouldNot(HaveOccurred())
					expectCart := model.Cart{
						Name: "aditira",
						Cart: []model.Product{
							{Id: "1", Name: "Tea", Price: "20000", Quantity: "3", Total: 60000},
						},
						TotalPrice: 60000,
					}
					Expect(cart).To(Equal(expectCart))
					userRepo.ResetUser()
					sessionRepo.ResetSessions()
					cartRepo.ResetCarts()
				})
			})

			When("add 1 product to cart with POST method and login cookie", func() {
				It("should added 1 cart data to file carts.json with `quantity per-product`, `price per-product` and `total price` calculated", func() {
					//register
					w := httptest.NewRecorder()
					r := httptest.NewRequest(http.MethodPost, "/user/register", strings.NewReader("username=aditira&password=12345"))
					r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					server.Handler().ServeHTTP(w, r)

					//login
					w = httptest.NewRecorder()
					r = httptest.NewRequest(http.MethodPost, "/user/login", strings.NewReader("username=aditira&password=12345"))
					r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					server.Handler().ServeHTTP(w, r)

					var cookie *http.Cookie
					var isCookieTokenExist bool
					for _, c := range w.Result().Cookies() {
						if c.Name == "session_token" {
							isCookieTokenExist = true
							cookie = c
						}
					}
					Expect(isCookieTokenExist).To(BeTrue())

					//add 1 product to cart
					w = httptest.NewRecorder()
					r = httptest.NewRequest(http.MethodPost, "/cart/add", strings.NewReader("product=1,Tea,20000,3"))
					r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					r.AddCookie(cookie)
					server.Handler().ServeHTTP(w, r)

					Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
					cart, err := cartRepo.ReadCart()
					Expect(err).ShouldNot(HaveOccurred())
					expectCart := model.Cart{
						Name: "aditira",
						Cart: []model.Product{
							{Id: "1", Name: "Tea", Price: "20000", Quantity: "3", Total: 60000},
						},
						TotalPrice: 60000,
					}
					Expect(cart).To(Equal(expectCart))
					userRepo.ResetUser()
					sessionRepo.ResetSessions()
					cartRepo.ResetCarts()
				})
			})

			When("add 3 product to cart with POST method and login cookie", func() {
				It("should added 3 cart data to file carts.json with `quantity per-product`, `price per-product` and `total price` calculated", func() {
					//register
					w := httptest.NewRecorder()
					r := httptest.NewRequest(http.MethodPost, "/user/register", strings.NewReader("username=aditira&password=12345"))
					r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					server.Handler().ServeHTTP(w, r)

					//login
					w = httptest.NewRecorder()
					r = httptest.NewRequest(http.MethodPost, "/user/login", strings.NewReader("username=aditira&password=12345"))
					r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					server.Handler().ServeHTTP(w, r)

					var cookie *http.Cookie
					var isCookieTokenExist bool
					for _, c := range w.Result().Cookies() {
						if c.Name == "session_token" {
							isCookieTokenExist = true
							cookie = c
						}
					}
					Expect(isCookieTokenExist).To(BeTrue())

					//add 3 product to cart
					w = httptest.NewRecorder()
					r = httptest.NewRequest(http.MethodPost, "/cart/add", strings.NewReader("product=1,Tea,20000,3&product=2,Milk,10000,4&product=6,Cake,4000,2"))
					r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					r.AddCookie(cookie)
					server.Handler().ServeHTTP(w, r)

					Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
					cart, err := cartRepo.ReadCart()
					Expect(err).ShouldNot(HaveOccurred())
					expectCart := model.Cart{
						Name: "aditira",
						Cart: []model.Product{
							{Id: "1", Name: "Tea", Price: "20000", Quantity: "3", Total: 60000},
							{Id: "2", Name: "Milk", Price: "10000", Quantity: "4", Total: 40000},
							{Id: "6", Name: "Cake", Price: "4000", Quantity: "2", Total: 8000},
						},
						TotalPrice: 108000,
					}
					Expect(cart).To(Equal(expectCart))
					userRepo.ResetUser()
					sessionRepo.ResetSessions()
					cartRepo.ResetCarts()
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
				It("should return logout success and token not found on database", func() {
					//register
					w := httptest.NewRecorder()
					r := httptest.NewRequest(http.MethodPost, "/user/register", strings.NewReader("username=aditira&password=12345"))
					r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					server.Handler().ServeHTTP(w, r)

					//login
					w = httptest.NewRecorder()
					r = httptest.NewRequest(http.MethodPost, "/user/login", strings.NewReader("username=aditira&password=12345"))
					r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					server.Handler().ServeHTTP(w, r)

					var cookie *http.Cookie
					var isCookieTokenExist bool
					for _, c := range w.Result().Cookies() {
						if c.Name == "session_token" {
							isCookieTokenExist = true
							cookie = c
						}
					}
					Expect(isCookieTokenExist).To(BeTrue())

					//logout
					w = httptest.NewRecorder()
					r = httptest.NewRequest(http.MethodGet, "/user/logout", nil)
					r.AddCookie(cookie)
					server.Handler().ServeHTTP(w, r)
					Expect(w.Result().StatusCode).To(Equal(http.StatusOK))

					listSession, _ := sessionRepo.ReadSessions()
					_, err := sessionRepo.TokenExist(listSession, cookie.Value)
					Expect(err.Error()).To(Equal("Token Not Found!"))

					userRepo.ResetUser()
				})
			})
		})

		Describe("/user/img/profile", func() {
			When("read image profile with POST method", func() {
				It("should return a wrong method", func() {
					w := httptest.NewRecorder()
					r := httptest.NewRequest(http.MethodPost, "/user/img/profile", nil)
					server.Handler().ServeHTTP(w, r)

					ErrorResponse := model.ErrorResponse{}
					json.NewDecoder(w.Body).Decode(&ErrorResponse)
					Expect(w.Result().StatusCode).To(Equal(http.StatusMethodNotAllowed))
					Expect(ErrorResponse.Error).To(Equal("Method is not allowed!"))
				})
			})

			When("read image profile with GET method, without login cookie", func() {
				It("should return a cookie not found", func() {
					w := httptest.NewRecorder()
					r := httptest.NewRequest(http.MethodGet, "/user/img/profile", nil)
					server.Handler().ServeHTTP(w, r)

					err := model.ErrorResponse{}
					json.NewDecoder(w.Body).Decode(&err)

					Expect(w.Result().StatusCode).To(Equal(http.StatusUnauthorized))
					Expect(err.Error).To(Equal("http: named cookie not present"))
				})
			})
		})

		Describe("/user/img/update-profile", func() {
			When("update image profile with GET method", func() {
				It("should return a wrong method", func() {
					w := httptest.NewRecorder()
					r := httptest.NewRequest(http.MethodGet, "/user/img/update-profile", nil)
					server.Handler().ServeHTTP(w, r)

					ErrorResponse := model.ErrorResponse{}
					json.NewDecoder(w.Body).Decode(&ErrorResponse)
					Expect(w.Result().StatusCode).To(Equal(http.StatusMethodNotAllowed))
					Expect(ErrorResponse.Error).To(Equal("Method is not allowed!"))
				})
			})

			When("update image profile with POST method, without login cookie", func() {
				It("should return a cookie not found", func() {
					w := httptest.NewRecorder()
					r := httptest.NewRequest(http.MethodPost, "/user/img/update-profile", nil)
					server.Handler().ServeHTTP(w, r)

					err := model.ErrorResponse{}
					json.NewDecoder(w.Body).Decode(&err)

					Expect(w.Result().StatusCode).To(Equal(http.StatusUnauthorized))
					Expect(err.Error).To(Equal("http: named cookie not present"))
				})
			})

			When("send empty request body with POST method", func() {
				It("should return internal server error", func() {
					//register
					w := httptest.NewRecorder()
					r := httptest.NewRequest(http.MethodPost, "/user/register", strings.NewReader("username=aditira&password=12345"))
					r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					server.Handler().ServeHTTP(w, r)

					//login
					w = httptest.NewRecorder()
					r = httptest.NewRequest(http.MethodPost, "/user/login", strings.NewReader("username=aditira&password=12345"))
					r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					server.Handler().ServeHTTP(w, r)

					var cookie *http.Cookie
					var isCookieTokenExist bool
					for _, c := range w.Result().Cookies() {
						if c.Name == "session_token" {
							isCookieTokenExist = true
							cookie = c
						}
					}
					Expect(isCookieTokenExist).To(BeTrue())

					//read image profile
					w = httptest.NewRecorder()
					r = httptest.NewRequest(http.MethodPost, "/user/img/update-profile", nil)
					r.AddCookie(cookie)
					server.Handler().ServeHTTP(w, r)

					Expect(w.Result().StatusCode).To(Equal(http.StatusInternalServerError))

					userRepo.ResetUser()
					sessionRepo.ResetSessions()
				})
			})

		})
	})

})
