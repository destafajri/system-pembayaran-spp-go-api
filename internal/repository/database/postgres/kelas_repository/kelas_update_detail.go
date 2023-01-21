package kelas_repository

import (
	"errors"
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
)

func (kelas *kelasImplementation) UpdateDetailKelas(kelasReq *entity.KelasEntity) (*model.UpdateDetailKelasResponse, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
		UPDATE kelas SET
			guru_id = $1,
			kelas =$2,
			updated_at = $3
		WHERE 
			id = $4
		`
	values := []interface{}{
		kelasReq.GuruID,
		kelasReq.Kelas,
		kelasReq.UpdatedAt,
		kelasReq.ID,
	}

	_, err := kelas.db.Exec(query, values...)
	if err != nil {
		log.Println(err)
		return nil, errors.New("error update detail kelas")
	}

	resp := model.UpdateDetailKelasResponse{
		ID: kelasReq.ID,
		GuruID: kelasReq.GuruID,
		Kelas: kelasReq.Kelas,
		Timestamp: kelasReq.Timestamp,
	}

	return &resp, nil
}