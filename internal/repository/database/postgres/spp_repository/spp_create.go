package spp_repository

import (
	"errors"
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
)

func (spp *sppImplementation) CreateSpp(request *entity.SppEntity) (*model.CreateSppResponse, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `INSERT INTO spp(
				id,
				siswa_id,
				no_spp,
				jatuh_tempo,
				jumlah,
				is_active,
				created_at,
				updated_at
			)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		`
	values := []interface{}{
		request.ID,
		request.SiswaID,
		request.NoSpp,
		request.JatuhTempo,
		request.Jumlah,
		request.IsActive,
		request.CreatedAt,
		request.UpdatedAt,
	}

	_, err := spp.db.Exec(query, values...)
	if err != nil {
		log.Println(err)
		return nil, errors.New("spp already exist")
	}

	resp := model.CreateSppResponse{
		ID:         request.ID,
		SiswaID:    request.SiswaID,
		JatuhTempo: request.JatuhTempo,
		Jumlah:     request.Jumlah,
		Status:     "unpaid",
		IsActive:   request.IsActive,
		Timestamp:  request.Timestamp,
	}

	return &resp, nil
}
