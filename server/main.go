package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"macromaster.com/packages/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal()
	}

	url := os.Getenv("DB_URL")


	test := api.CreateServer(":4000", url)
	test.StartServer()
}