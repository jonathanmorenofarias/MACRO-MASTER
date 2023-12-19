package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_"github.com/lib/pq"
)

func CreateServer(port string, url string) *Server {
	return (
		&Server {
			Port: port,
			DB_URL: url,
	})
}

func (api Server) StartServer(){
	router := mux.NewRouter()

	fmt.Printf("Server is listening on port: %s\n", api.Port)

	api.DB = api.StartDB()
	defer api.DB.Close()

	router.HandleFunc("/signup", printMethod(api.handleRegister))
	router.HandleFunc("/login", printMethod(api.handleLogin))
	router.HandleFunc("/test", AuthenticateUser(func (w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "Im done")
	}))

	http.ListenAndServe(api.Port, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	)(router))
}

func (api Server) StartDB () *sql.DB {
	db, err := sql.Open("postgres", api.DB_URL)

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	query := `CREATE TABLE IF NOT EXISTS "Users" (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL UNIQUE,
		username VARCHAR(100) NOT NULL UNIQUE,
		password VARCHAR(100) NOT NULL,
		created timestamptz DEFAULT CURRENT_TIMESTAMP
	)`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to DB!")


	return db
}

func (api Server) RespondStatus(w http.ResponseWriter, status int, message any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(message)
}

func (api Server) handleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var person User

		err := json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			api.RespondStatus(w, 400, ReqError{Error: "Access Denied"})
			return
		}

		if person.Name == "" || person.Email == "" || person.Username == "" || person.Password == "" {
			api.RespondStatus(w, 400, ReqError{Error: "Access Denied"})
			return
		}

		hashedPassowrd := HashPassword(person.Password)

		query := `INSERT INTO "Users" (name, email, username, password)
			VALUES ($1, $2, $3, $4) RETURNING id`

		var id int
		err = api.DB.QueryRow(query, person.Name, person.Email, person.Username, string(hashedPassowrd)).Scan(&id)

		if err != nil {
			api.RespondStatus(w, 409, ReqError{Error: "Username or Password already exists"})
		}

		createdUser := fmt.Sprintf(`{"id": "%d",  "name": "%s", "username": "%s" }`, id, person.Name, person.Username)
		api.RespondStatus(w, 200, ReqSuccess{Success: createdUser})
	}
}

func (api Server) handleLogin(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		var person User
		err := json.NewDecoder(r.Body).Decode(&person)

		if err != nil {
			api.RespondStatus(w, 400, ReqError{Error: "Access Denied"})
			return
		}

		if person.Username == "" {
			api.RespondStatus(w, 400, ReqError{Error: "Access Denied"})
			return
		} else if person.Password == "" {
			api.RespondStatus(w, 400, ReqError{Error: "Access Denied"})
			return
		}

		query := `SELECT id, password 
				FROM "Users"
				WHERE username = $1`

		var verifyPerson UserAuth
		err = api.DB.QueryRow(query, person.Username).Scan(&verifyPerson.Id, &verifyPerson.Password)
		
		if err != nil {
			api.RespondStatus(w, 404, ReqError{Error: "Access Denied"})
			return
		}

		match := DecryptPassword(person.Password, verifyPerson.Password)

		if !match {
			api.RespondStatus(w, 401, ReqError{Error: "Access Denied"})
		} else {
			token, err := CreateToken(verifyPerson.Id)
			if err != nil {
				api.RespondStatus(w, 500, ReqError{Error: "Server Error"})
				return
			}
			api.RespondStatus(w, 200, token)
		}
	}
}