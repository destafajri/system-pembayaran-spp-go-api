package model

import (
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/exception"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/entity"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateSppRequest struct {
	SiswaNIS   int    `json:"siswa_nis"`
	NoSpp      string `json:"no_spp"`
	JatuhTempo string `json:"jatuh_tempo"`
	Jumlah     string `json:"jumlah"`
}

type CreateSppResponse struct {
	ID         string `json:"id"`
	SiswaID    string `json:"siswa_id"`
	JatuhTempo string `json:"jatuh_tempo"`
	Jumlah     string `json:"jumlah"`
	Status     string `json:"status"`
	IsActive   bool   `json:"is_active"`
	entity.Timestamp
}

// validation
func ValidateCreateSppInput(request *CreateSppRequest) error {
	err := validation.ValidateStruct(request,
		validation.Field(&request.SiswaNIS, validation.Required),
		validation.Field(&request.NoSpp, validation.Required),
		validation.Field(&request.Jumlah, validation.Required),
		validation.Field(&request.JatuhTempo, validation.Required),
	)

	if err != nil {
		log.Println(err)
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}

	return nil
}
