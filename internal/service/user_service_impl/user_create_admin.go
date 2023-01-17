package user_service

import (
	"log"
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/google/uuid"
)

func (user *userServiceimpl) CreateAdmin(request *model.CreateAdminRequest, timestamp time.Time) (*model.CreateAdminResponse, error) {
	//validation input
	err := model.ValidateCreateUserInput(request)
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
		Password:  request.Password,
		Role:      "admin",
		IsActive:  true,
		Timestamp: t,
	}

	resp, err := user.userRepository.CreateAdmin(&input)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return resp, nil
}
