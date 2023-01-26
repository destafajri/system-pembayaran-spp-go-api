package spp_service

import (
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/validations"
	"github.com/google/uuid"
)

func (spp *sppServiceimpl) CreateSpp(request *model.CreateSppRequest, timestamp time.Time) (*model.CreateSppResponse, error) {
	// validation input
	err := model.ValidateCreateSppInput(request)
	if err != nil {
		return nil, err
	}

	// validation currency
	ok, err := validations.IDRValidation(request.Jumlah)
	if err != nil && !ok {
		return nil, err
	}

	// validation date
	date, err := validations.DateValidation(request.JatuhTempo)
	if err != nil && !date {
		return nil, err
	}

	user_id, err := spp.sppRepository.GetSiswaID(request.SiswaNIS)
	if err != nil {
		return nil, err
	}

	t := entity.Timestamp{
		CreatedAt: timestamp.Local(),
		UpdatedAt: timestamp.Local(),
	}

	req := entity.SppEntity{
		ID:         uuid.New().String(),
		SiswaID:    user_id,
		NoSpp:      request.NoSpp,
		JatuhTempo: request.JatuhTempo,
		Jumlah:     request.Jumlah,
		IsActive:   true,
		Timestamp:  t,
	}

	resp, err := spp.sppRepository.CreateSpp(&req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
