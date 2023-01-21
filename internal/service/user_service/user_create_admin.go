package user_service

import (
	"log"
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
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

	t := entity.Timestamp{
		CreatedAt: timestamp.Local(),
		UpdatedAt: timestamp.Local(),
	}

	input := entity.UserEntity{
		ID:        uuid.New().String(),
		Email:     request.Email,
		Username:  request.Username,
		Password:  request.Password,
		Role:      "admin",
		IsActive:  true,
		Timestamp: t,
	}

	users, err := user.userRepository.CreateAdmin(&input)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp := model.CreateAdminResponse{
		ID:        users.ID,
		Email:     users.Email,
		Username:  users.Username,
		Role:      users.Role,
		IsActive:  users.IsActive,
		Timestamp: users.Timestamp,
	}

	return &resp, nil
}
