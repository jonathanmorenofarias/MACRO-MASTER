package api

import (
	"log"
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"

)

func CreateToken(id int) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(10 * time.Minute).Unix()
	claims["id"] = id

	secret := os.Getenv("JWT_SECRET")

	signedToken, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Fatal(err)
	}

	return signedToken
}