package kelas_service

import (
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/validations"
	"github.com/pkg/errors"
)

func (kelas *kelasServiceimpl) DeleteKelas(kelas_id string) error {
	if err := validations.ValidateUUID(kelas_id); err != nil {
		log.Println(err)
		return err
	}

	// check if user id found
	kelasfound, err := kelas.kelasRepository.CekKelasExistByID(kelas_id)
	if err != nil && !kelasfound {
		log.Println(err)
		return err
	}

	err = kelas.kelasRepository.DeleteKelas(kelas_id)
	if err != nil {
		log.Println(err)
		return errors.Wrap(err, "delete kelas on service")
	}

	return nil
}
