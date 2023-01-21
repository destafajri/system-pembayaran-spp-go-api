package model

import (
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/exception"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateKelasRequest struct {
	GuruID string `json:"guru_id"`
	Kelas  string `json:"kelas"`
}

type CreateKelasResponse struct {
	ID     string `json:"id"`
	GuruID string `json:"guru_id"`
	Kelas  string `json:"kelas"`
}

type GetListKelasResponse struct {
	ID        string `json:"id"`
	GuruID    string `json:"guru_id"`
	WaliKelas string `json:"walikelas"`
	Kelas     string `json:"kelas"`
}

// validation
func ValidateCreateKelasInput(request *CreateKelasRequest) error {
	err := validation.ValidateStruct(request,
		validation.Field(&request.GuruID, validation.Required),
		validation.Field(&request.Kelas, validation.Required),
	)

	if err != nil {
		log.Println(err)
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}

	return nil
}
