package api

import (
	"database/sql"
	"fmt"
	"net/http"
)

func printMethod(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s\n", r.Method, r.URL)
		handler(w, r)
	}
}

type Server struct {
	Port   string
	DB_URL string
	DB     *sql.DB
}

type ReqError struct {
	Error string
}

type ReqSuccess struct {
	Success string
}

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserAuth struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
}
