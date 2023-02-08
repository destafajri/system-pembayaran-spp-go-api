package guru_service

import (
	"log"
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/validations"
	"github.com/google/uuid"
)

func (guru *guruServiceimpl) CreateGuru(request *model.CreateGuruRequest, timestamp time.Time) (*model.CreateGuruResponse, error) {
	// validation input
	err := model.ValidateCreateGuruInput(request)
	if err != nil {
		log.Println()
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

	input1 := entity.UserEntity{
		ID:        uuid.New().String(),
		Email:     request.Email,
		Username:  request.Username,
		Password:  request.Password,
		Role:      "guru",
		IsActive:  true,
		Timestamp: t,
	}

	input2 := entity.GuruEntity{
		ID:       uuid.New().String(),
		UserID:   input1.ID,
		Nama:     request.Nama,
		IsActive: true,
	}

	resp, err := guru.guruRepository.CreateGuru(&input1, &input2)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return resp, nil
}
