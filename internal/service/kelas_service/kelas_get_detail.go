package kelas_service

import (
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/validations"
	"github.com/pkg/errors"
)

func (kelas *kelasServiceimpl) GetDetailKelas(kelas_id string) (*model.GetDetailKelasResponse, error) {
	if err := validations.ValidateUUID(kelas_id); err != nil {
		log.Println(err)
		return nil, err
	}

	// check if kelas id found
	kelasfound, err := kelas.kelasRepository.CekKelasExistByID(kelas_id)
	if err != nil && !kelasfound {
		log.Println(err)
		return nil, err
	}

	kelasDetail, err := kelas.kelasRepository.GetDetailKelas(kelas_id)
	if err != nil {
		log.Println(err)
		return nil, errors.Wrap(err, "find kelas by kelas Id on service")
	}

	return kelasDetail, nil
}