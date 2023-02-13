package password

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	//hash pass using bcrypt
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		err = fmt.Errorf("decoding request body: %w", err)
		log.Println(err)
		return "", err
	}

	return string(hashPassword), nil
}
