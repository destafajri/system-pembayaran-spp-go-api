package user_service

import (
	"log"

	"github.com/pkg/errors"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	validations "github.com/destafajri/system-pembayaran-spp-go-api/internal/validation"
)

func (user *userServiceimpl) GetDetailUser(user_id string) (*model.GetDetailUser, error) {
	if err := validations.ValidateUUID(user_id); err != nil {
		log.Println(err)
		return nil, err
	}

	// check if user id found
	userfound, err := user.userRepository.CekUserExistByID(user_id)
	if err != nil && !userfound {
		log.Println(err)
		return nil, err
	}

	userDetail, err := user.userRepository.GetDetailUser(user_id)
	if err != nil {
		log.Println(err)
		return nil, errors.Wrap(err, "find user by user Id on service")
	}

	return userDetail, nil
}
