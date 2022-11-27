package auth

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// EcnryptionPassword
// Encrypt the password text
func EncryptionPassword(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error encryption password: ", err.Error())
		return ""
	}

	return string(hashed)
}

// Decrypt
// decrypt the password hashing
// and return true and false
func DecryptionPassword(hashed, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		return false
	} else {
		return true
	}
}
