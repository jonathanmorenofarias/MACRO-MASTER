package api

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"encoding/json"
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

	query := `CREATE TABLE IF NOT EXISTS "user" (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL,
		username VARCHAR(100) NOT NULL,
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

	http.ListenAndServe(api.Port, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	)(router))

}

func (api Server) handleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var person User

		json.NewDecoder(r.Body).Decode(&person)

		query := `INSERT INTO "user" (name, email, username, password)
			VALUES ($1, $2, $3, $4) RETURNING id`

		var key int
		err := api.DB.QueryRow(query, person.Name, person.Email, person.Username, person.Password).Scan(&key)

		if err != nil {
			log.Fatal()
		}

		fmt.Fprintf(w, "Account for id: %d created successfully.", key)
	}
}