package password

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CompareHashPassword(passwordPayload, passwordFromDB string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(passwordFromDB), []byte(passwordPayload)); err != nil {
		log.Println(err)
		return errors.New("wrong password")
	}

	return nil
}
