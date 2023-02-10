package spp_service

import (
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/validations"
	"github.com/pkg/errors"
)

func (spp *sppServiceimpl) DeleteSpp(spp_id string) error {
	if err := validations.ValidateUUID(spp_id); err != nil {
		log.Println(err)
		return err
	}

	// check if user id found
	sppfound, err := spp.sppRepository.CekSppExistByID(spp_id)
	if err != nil && !sppfound {
		log.Println(err)
		return err
	}

	// cek status spp
	status, _ := spp.sppRepository.CekStatusBayar(spp_id)

	if status == "paid" {
		err = spp.sppRepository.DeleteStatus(spp_id)
		if err != nil {
			return err
		}
	}

	err = spp.sppRepository.DeleteSpp(spp_id)
	if err != nil {
		log.Println(err)
		return errors.Wrap(err, "delete spp on service")
	}

	return nil
}
