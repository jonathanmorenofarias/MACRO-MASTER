package api

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var secret string = os.Getenv("JWT_SECRET")


func CreateToken(id int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(10 * time.Minute).Unix()
	claims["id"] = id

	signedToken, err := token.SignedString([]byte(secret))

	return signedToken, err
}

func AuthenticateUser(handler http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		tokenString := r.Header.Get("token")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
	
			return []byte(secret), nil
		})

		if err != nil {
			fmt.Println(err)
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		fmt.Println(claims["id"])

		handler(w, r)
	}
}