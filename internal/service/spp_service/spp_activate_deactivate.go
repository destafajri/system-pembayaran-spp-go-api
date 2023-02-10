package spp_service

import (
	"errors"
	"log"
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/validations"
)

func (spp *sppServiceimpl) ActivateSpp(spp_id string, timestamp time.Time) error{
	if err := validations.ValidateUUID(spp_id); err != nil {
		log.Println(err)
		return err
	}

	// check if spp id found
	sppfound, err := spp.sppRepository.CekSppExistByID(spp_id)
	if err != nil && !sppfound {
		log.Println(err)
		return err
	}

	err = spp.sppRepository.ActivateSpp(spp_id, timestamp)
	if err != nil {
		log.Println(err)
		return errors.New("activate spp by spp Id on service")
	}

	return nil
}

func (spp *sppServiceimpl) DeactivateSpp(spp_id string, timestamp time.Time) error{
	if err := validations.ValidateUUID(spp_id); err != nil {
		log.Println(err)
		return err
	}

	// check if spp id found
	sppfound, err := spp.sppRepository.CekSppExistByID(spp_id)
	if err != nil && !sppfound {
		log.Println(err)
		return err
	}

	err = spp.sppRepository.DeactivateSpp(spp_id, timestamp)
	if err != nil {
		log.Println(err)
		return errors.New("deactivate spp by spp Id on service")
	}

	return nil
}