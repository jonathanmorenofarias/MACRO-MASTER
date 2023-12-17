package api

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string) {
	hashedPassowrd, _ := bcrypt.GenerateFromPassword([]byte(password), 12)

	return string(hashedPassowrd)
}

func DecryptPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}