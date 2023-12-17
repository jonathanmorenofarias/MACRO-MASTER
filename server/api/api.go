package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
)


func printMethod(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s\n", r.Method, r.URL)
		handler.ServeHTTP(w, r)
	}
}

func CreateServer(port string, url string) *Server {
	return (
		&Server {
			Port: port,
			DB_URL: url,
	})
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

func (api Server) StartServer(){
	router := mux.NewRouter()

	fmt.Printf("Server is listening on port: %s\n", api.Port)

	api.DB = api.StartDB()
	defer api.DB.Close()

	router.HandleFunc("/signup", printMethod(api.handleRegister))
	router.HandleFunc("/login", printMethod(api.handleLogin))

	http.ListenAndServe(api.Port, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	)(router))
}

func (api Server) handleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "text/plain")

		var person User

		err := json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			http.Error(w, "Could not parse JSON", http.StatusBadRequest)
			return
		}

		if person.Name == "" {
			http.Error(w, "Missing name", http.StatusBadRequest)
			return
		} else if person.Email == "" {
			http.Error(w, "Missing email", http.StatusBadRequest)
			return
		} else if person.Username == "" {
			http.Error(w, "Missing username", http.StatusBadRequest)
			return
		} else if person.Password == "" {
			http.Error(w, "Missing password", http.StatusBadRequest)
			return
		}

		hashedPassowrd := HashPassword(person.Password)

		query := `INSERT INTO "Users" (name, email, username, password)
			VALUES ($1, $2, $3, $4) RETURNING id`

		var key int
		err = api.DB.QueryRow(query, person.Name, person.Email, person.Username, string(hashedPassowrd)).Scan(&key)

		if err != nil {
			pgErr, ok := err.(*pq.Error)
			if ok {
				if pgErr.Code == "23505" {
					http.Error(w, "This username or email already exists!", http.StatusBadRequest)
					return
				}
			}
		}

		fmt.Fprintf(w, "Account for id: %d created successfully.", key)
	}
}

func (api Server) handleLogin(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "text/plain")

		var person User

		err := json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			http.Error(w, "Could not parse JSON", http.StatusBadRequest)
			return
		}

		if person.Username == "" {
			http.Error(w, "Missing username", http.StatusBadRequest)
			return
		} else if person.Password == "" {
			http.Error(w, "Missing password", http.StatusBadRequest)
			return
		}

		query := `SELECT password FROM "Users"
			WHERE username = $1`

		var password string
		err = api.DB.QueryRow(query, person.Username).Scan(&password)

		if err != nil {
			log.Fatal(err)
		}

		match := DecryptPassword(person.Password, password)
		if match {
			fmt.Println("it is a match")
		} else {
			fmt.Println("not a match")
		}

	}
}