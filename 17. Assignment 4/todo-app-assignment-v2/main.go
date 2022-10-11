package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/helper"
	"a21hc3NpZ25tZW50/middleware"
	"a21hc3NpZ25tZW50/model"

	"github.com/google/uuid"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var credentials model.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)

	if err == io.EOF {
		helper.GenerateResponseError(w, r, 400, "Internal Server Error")
		return
	}

	if credentials.Username == "" || credentials.Password == "" {
		helper.GenerateResponseError(w, r, 400, "Username or Password empty")
		return
	}

	for username := range db.Users {
		if credentials.Username == username {
			helper.GenerateResponseError(w, r, 409, "Username already exist")

			return
		}
	}

	db.Users[credentials.Username] = credentials.Password

	successResponse := model.SuccessResponse{
		Username: credentials.Username,
		Message:  "Register Success",
	}

	db.Task[credentials.Username] = []model.Todo{}

	helper.GenerateResponseSuccess(w, r, 200, successResponse)
	// TODO: answer here
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials model.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)

	if err == io.EOF {
		helper.GenerateResponseError(w, r, 400, "Internal Server Error")
		return
	}

	if credentials.Username == "" || credentials.Password == "" {
		helper.GenerateResponseError(w, r, 400, "Username or Password empty")
		return
	}

	for username, password := range db.Users {
		if credentials.Username == username && credentials.Password == password {
			sessiont_token := uuid.New()

			newSession := model.Session{
				Username: credentials.Username,
				Expiry:   time.Now().Local().Add(time.Hour * 5),
			}

			// SET COOKIES in Client
			c := &http.Cookie{}
			c.Name = "session_token"
			c.Value = sessiont_token.String()
			c.Expires = time.Now().Add(5 * time.Hour)

			http.SetCookie(w, c)

			// SET COOKIES in memory
			db.Sessions[sessiont_token.String()] = newSession

			successResponse := model.SuccessResponse{
				Username: credentials.Username,
				Message:  "Login Success",
			}

			helper.GenerateResponseSuccess(w, r, 200, successResponse)

			return
		}
	}

	helper.GenerateResponseError(w, r, 401, "Wrong User or Password!")
	// TODO: answer here
}

func AddToDo(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	var todo model.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)

	if err == io.EOF {
		helper.GenerateResponseError(w, r, 400, "Internal Server Error")
		return
	}

	id := uuid.New()

	todo.Id = id.String()
	coockie, _ := r.Cookie("session_token")

	db.Task[db.Sessions[coockie.Value].Username] = append(db.Task[db.Sessions[coockie.Value].Username], todo)

	successResponse := model.SuccessResponse{
		Username: db.Sessions[coockie.Value].Username,
		Message:  "Task Create a todo app program added!",
	}

	helper.GenerateResponseSuccess(w, r, 200, successResponse)
}

func ListToDo(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	coockie, _ := r.Cookie("session_token")
	username := db.Sessions[coockie.Value].Username
	_, key := db.Task[username]

	if !key {
		helper.GenerateResponseError(w, r, http.StatusNotFound, "Todolist not found!")
		return
	} else if db.Task[username] == nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(make([]string, 0))
		return
	}

	foo_marshalled, err := json.Marshal(db.Task[username])

	if err != nil {
		helper.GenerateResponseError(w, r, 500, "Internal Server Error")
		return
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(foo_marshalled))
}

func ClearToDo(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	coockie, _ := r.Cookie("session_token")
	username := db.Sessions[coockie.Value].Username

	db.Task[username] = []model.Todo{}

	successResponse := model.SuccessResponse{
		Username: username,
		Message:  "Clear ToDo Success",
	}

	helper.GenerateResponseSuccess(w, r, 200, successResponse)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	coockie, _ := r.Cookie("session_token")

	successResponse := model.SuccessResponse{
		Username: db.Sessions[coockie.Value].Username,
		Message:  "Logout Success",
	}

	delete(db.Sessions, coockie.Value)

	helper.GenerateResponseSuccess(w, r, 200, successResponse)
	// TODO: answer here
}

func ResetToDo(w http.ResponseWriter, r *http.Request) {
	db.Task = map[string][]model.Todo{}
	w.WriteHeader(http.StatusOK)
}

type API struct {
	mux *http.ServeMux
}

func NewAPI() API {
	mux := http.NewServeMux()
	api := API{
		mux,
	}

	mux.Handle("/user/register", middleware.Post(http.HandlerFunc(Register)))
	mux.Handle("/user/login", middleware.Post(http.HandlerFunc(Login)))
	mux.Handle("/user/logout", middleware.Get(middleware.Auth(http.HandlerFunc(Logout))))

	// TODO: answer here
	mux.Handle("/todo/create", middleware.Post(middleware.Auth(http.HandlerFunc(AddToDo))))
	mux.Handle("/todo/read", middleware.Get(middleware.Auth(http.HandlerFunc(ListToDo))))
	mux.Handle("/todo/clear", middleware.Delete(middleware.Auth(http.HandlerFunc(ClearToDo))))

	mux.Handle("/todo/reset", http.HandlerFunc(ResetToDo))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}

func main() {
	mainAPI := NewAPI()
	mainAPI.Start()
}
