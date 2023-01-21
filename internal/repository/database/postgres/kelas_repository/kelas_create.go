package kelas_repository

import (
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/pkg/errors"
)

func (kelas *kelasImplementation) CreateKelas(kelasReq *entity.KelasEntity) (*model.CreateKelasResponse, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `INSERT INTO kelas(
			id,
			guru_id,
			kelas
		)
		VALUES ($1, $2, $3)
	`
	values := []interface{}{
		kelasReq.ID,
		kelasReq.GuruID,
		kelasReq.Kelas,
	}

	_, err := kelas.db.Exec(query, values...)
	if err != nil {
		log.Println(err)
		return nil, errors.New("kelas already exist")
	}

	resp := model.CreateKelasResponse{
		ID:     kelasReq.ID,
		GuruID: kelasReq.GuruID,
		Kelas:  kelasReq.Kelas,
	}

	return &resp, nil
}
