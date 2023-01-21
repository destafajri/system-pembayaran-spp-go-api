package model

import (
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/exception"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/entity"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateGuruRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Nama     string `json:"nama"`
}

type CreateGuruResponse struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Nama     string `json:"nama"`
	Role     string `json:"role"`
	IsActive bool   `json:"is_active"`
	entity.Timestamp
}

type GetListGuruResponse struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Nama     string `json:"nama"`
	Role     string `json:"role"`
	IsActive bool   `json:"is_active"`
	entity.Timestamp
}

type GetDetailGuruResponse struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Nama     string `json:"nama"`
	Role     string `json:"role"`
	IsActive bool   `json:"is_active"`
	entity.Timestamp
}

// validation
func ValidateCreateGuruInput(request *CreateGuruRequest) error {
	err := validation.ValidateStruct(request,
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.Password, validation.Required),
		validation.Field(&request.Username, validation.Required),
		validation.Field(&request.Nama, validation.Required),
	)

	if err != nil {
		log.Println(err)
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}

	return nil
}
