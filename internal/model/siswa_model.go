package model

import (
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/exception"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/entity"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateSiswaRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Kelas    string `json:"kelas"`
	NIS      int    `json:"nis"`
	Nama     string `json:"nama"`
	Angkatan string `json:"angkatan"`
}

type CreateSiswaResponse struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	KelasID  string `json:"kelas_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     string `json:"role"`
	NIS      int    `json:"nis"`
	Nama     string `json:"nama"`
	Angkatan string `json:"angkatan"`
	IsActive bool   `json:"is_active"`
	entity.Timestamp
}

type GetListSiswaResponse struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	KelasID  string `json:"kelas_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     string `json:"role"`
	NIS      int    `json:"nis"`
	Nama     string `json:"nama"`
	Kelas    string `json:"kelas"`
	Angkatan string `json:"angkatan"`
	IsActive bool   `json:"is_active"`
	entity.Timestamp
}

// validation
func ValidateCreateSiswaInput(request *CreateSiswaRequest) error {
	err := validation.ValidateStruct(request,
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.Username, validation.Required),
		validation.Field(&request.Password, validation.Required),
		validation.Field(&request.Kelas, validation.Required),
		validation.Field(&request.NIS, validation.Required),
		validation.Field(&request.Nama, validation.Required),
		validation.Field(&request.Angkatan, validation.Required),
	)

	if err != nil {
		log.Println(err)
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}

	return nil
}
