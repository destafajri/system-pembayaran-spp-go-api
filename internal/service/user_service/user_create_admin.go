package user_service

import (
	"log"
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/helper/password"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"
	validations "github.com/destafajri/system-pembayaran-spp-go-api/internal/validations"
	"github.com/google/uuid"
)

func (user *userServiceimpl) CreateAdmin(request *model.CreateAdminRequest, timestamp time.Time) (*model.CreateAdminResponse, error) {
	// validation input
	err := model.ValidateCreateUserInput(request)
	if err != nil {
		return nil, err
	}

	// validation
	email, err := validations.EmailValidation(request.Email)
	if err != nil && !email {
		return nil, err
	}

	username, err := validations.UsernameValidation(request.Username)
	if err != nil && !username {
		return nil, err
	}

	pass, err := validations.PasswordValidation(request.Password)
	if err != nil && !pass {
		return nil, err
	}

	// hash password
	passwordHash, err := password.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	t := entity.Timestamp{
		CreatedAt: timestamp.Local(),
		UpdatedAt: timestamp.Local(),
	}

	input := entity.UserEntity{
		ID:        uuid.New().String(),
		Email:     request.Email,
		Username:  request.Username,
		Password:  passwordHash,
		Role:      "admin",
		IsActive:  true,
		Timestamp: t,
	}

	users, err := user.userRepository.CreateAdmin(&input)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return users, nil
}
