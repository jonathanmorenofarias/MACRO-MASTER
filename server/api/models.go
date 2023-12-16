package api

import (
	"database/sql"
)

type Server struct {
	Port string
	DB_URL string
	DB *sql.DB
}

type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Confirm string `json:"confirm"`
}