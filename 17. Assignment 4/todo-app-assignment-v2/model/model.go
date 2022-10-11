package model

import "time"

type Todo struct {
	Id   string `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

type Session struct {
	Username string
	Expiry   time.Time
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}
