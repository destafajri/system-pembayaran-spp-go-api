package model

import (
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/exception"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type BayarSppRequest struct {
	NoSpp string `json:"no_spp"`
}

type BayarSppResponse struct {
	ID         string `json:"id"`
	SiswaID    string `json:"siswa_id"`
	NoSpp      string `json:"no_spp"`
	JatuhTempo string `json:"jatuh_tempo"`
	Jumlah     string `json:"jumlah"`
	Status     string `json:"status"`
}

func ValidateCallbackBayarInput(request *BayarSppRequest) error {
	err := validation.ValidateStruct(request,
		validation.Field(&request.NoSpp, validation.Required),
	)

	if err != nil {
		log.Println(err)
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}

	return nil
}
