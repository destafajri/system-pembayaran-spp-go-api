package guru_service

import (
	"log"
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/validations"
	"github.com/pkg/errors"
)

func (guru *guruServiceimpl) ActivateGuru(guru_id string, timestamp time.Time) error{
	if err := validations.ValidateUUID(guru_id); err != nil {
		log.Println(err)
		return err
	}

	// check if guru id found
	userfound, err := guru.guruRepository.CekGuruExistByID(guru_id)
	if err != nil && !userfound {
		log.Println(err)
		return err
	}

	// get user id
	user_id, err := guru.guruRepository.GetUserID(guru_id)
	if err != nil {
		log.Println(err)
		return errors.Wrap(err, "getting user id by guru id")
	}

	err = guru.guruRepository.ActivateGuru(user_id, guru_id, timestamp)
	if err != nil {
		log.Println(err)
		return errors.Wrap(err, "activate guru by guru Id on service")
	}

	return nil
}

func (guru *guruServiceimpl) DeactivateGuru(guru_id string, timestamp time.Time) error{
	if err := validations.ValidateUUID(guru_id); err != nil {
		log.Println(err)
		return err
	}

	// check if user id found
	userfound, err := guru.guruRepository.CekGuruExistByID(guru_id)
	if err != nil && !userfound {
		log.Println(err)
		return err
	}

	// get user id
	user_id, err := guru.guruRepository.GetUserID(guru_id)
	if err != nil {
		log.Println(err)
		return errors.Wrap(err, "getting user id by guru id")
	}

	err = guru.guruRepository.DeactivateGuru(user_id, guru_id, timestamp)
	if err != nil {
		log.Println(err)
		return errors.Wrap(err, "deactivate guru by guru Id on service")
	}
	
	return nil
}