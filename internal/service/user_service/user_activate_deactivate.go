package user_service

import (
	"log"
	"time"

	validations "github.com/destafajri/system-pembayaran-spp-go-api/internal/validations"
	"github.com/pkg/errors"
)

func (user *userServiceimpl) ActivateUser(id string, timestamp time.Time) error{
	if err := validations.ValidateUUID(id); err != nil {
		log.Println(err)
		return err
	}

	// check if user id found
	userfound, err := user.userRepository.CekUserExistByID(id)
	if err != nil && !userfound {
		log.Println(err)
		return err
	}

	err = user.userRepository.ActivateUser(id, timestamp)
	if err != nil {
		log.Println(err)
		return errors.Wrap(err, "activate user by user Id on service")
	}

	return nil
}

func (user *userServiceimpl) DeactivateUser(id string, timestamp time.Time) error{
	if err := validations.ValidateUUID(id); err != nil {
		log.Println(err)
		return err
	}

	// check if user id found
	userfound, err := user.userRepository.CekUserExistByID(id)
	if err != nil && !userfound {
		log.Println(err)
		return err
	}

	err = user.userRepository.DeactivateUser(id, timestamp)
	if err != nil {
		log.Println(err)
		return errors.Wrap(err, "deactivate user by user Id on service")
	}
	
	return nil
}