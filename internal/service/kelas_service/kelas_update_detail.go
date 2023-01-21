package kelas_service

import (
	"log"
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/validations"
	"github.com/pkg/errors"
)

func (kelas *kelasServiceimpl) UpdateDetailKelas(kelas_id string, kelasReq *model.UpdateDetailKelasRequest) (*model.UpdateDetailKelasResponse, error) {
	if err := validations.ValidateUUID(kelas_id); err != nil {
		log.Println(err)
		return nil, err
	}

	// input validation
	err := model.ValidateUpdateKelasInput(kelasReq)
	if err != nil {
		return nil, err
	}

	// check if kelas id found
	userfound, err := kelas.kelasRepository.CekKelasExistByID(kelas_id)
	if err != nil && !userfound {
		log.Println(err)
		return nil, err
	}

	t := entity.Timestamp{
		UpdatedAt: time.Now().Local(),
	}

	req := entity.KelasEntity{
		ID:        kelas_id,
		GuruID:    kelasReq.GuruID,
		Kelas:     kelasReq.Kelas,
		Timestamp: t,
	}

	resp, err := kelas.kelasRepository.UpdateDetailKelas(&req)
	if err != nil {
		log.Println(err)
		return nil, errors.Wrap(err, "update kelas on service")
	}

	return resp, nil
}
