package kelas_service

import (
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/google/uuid"
)

func (kelas *kelasServiceimpl) CreateKelas(request *model.CreateKelasRequest) (*model.CreateKelasResponse, error) {
	// validation input
	err := model.ValidateCreateKelasInput(request)
	if err != nil {
		return nil, err
	}

	// check if guru id found
	gurufound, err := kelas.kelasRepository.CekGuruExistByID(request.GuruID)
	if err != nil && !gurufound {
		log.Println(err)
		return nil, err
	}

	input := entity.KelasEntity{
		ID:     uuid.New().String(),
		GuruID: request.GuruID,
		Kelas:  request.Kelas,
	}

	resp, err := kelas.kelasRepository.CreateKelas(&input)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return resp, nil
}
