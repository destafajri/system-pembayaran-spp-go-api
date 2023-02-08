package siswa_service

import (
	"errors"
	_ "errors"
	"log"
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/validations"
	"github.com/google/uuid"
)

func (siswa *siswaServiceimpl) CreateSiswa(request *model.CreateSiswaRequest, timestamp time.Time) (*model.CreateSiswaResponse, error) {
	// validation input
	err := model.ValidateCreateSiswaInput(request)
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

	// cek nim
	nimFound , _ := siswa.siswaRepository.CekNIS(request.NIS)
	if nimFound != "" {
		return nil, errors.New("NIM already exist")
	}

	t := entity.Timestamp{
		CreatedAt: timestamp.Local(),
		UpdatedAt: timestamp.Local(),
	}

	// get kelas id
	kelas_id, err := siswa.siswaRepository.GetKelasID(request.Kelas)
	if err != nil {
		return nil, err
	}

	input1 := entity.UserEntity{
		ID:        uuid.New().String(),
		Email:     request.Email,
		Username:  request.Username,
		Password:  request.Password,
		Role:      "siswa",
		IsActive:  true,
		Timestamp: t,
	}

	input2 := entity.SiswaEntity{
		ID:       uuid.New().String(),
		UserID:   input1.ID,
		KelasID:  kelas_id,
		NIS:      request.NIS,
		Nama:     request.Nama,
		Angkatan: request.Angkatan,
		IsActive: true,
	}

	resp, err := siswa.siswaRepository.CreateSiswa(&input1, &input2)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return resp, nil
}