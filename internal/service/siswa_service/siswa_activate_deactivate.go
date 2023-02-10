package siswa_service

import (
	"log"
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/validations"
	"github.com/pkg/errors"
)

func (siswa *siswaServiceimpl) ActivateSiswa(siswa_id string, timestamp time.Time) error{
	if err := validations.ValidateUUID(siswa_id); err != nil {
		log.Println(err)
		return err
	}

	// check if siswa id found
	userfound, err := siswa.siswaRepository.CekSiswaExistByID(siswa_id)
	if err != nil && !userfound {
		log.Println(err)
		return err
	}

	// get user id
	user_id, err := siswa.siswaRepository.GetUserID(siswa_id)
	if err != nil {
		log.Println(err)
		return errors.Wrap(err, "getting user id by siswa id")
	}

	err = siswa.siswaRepository.ActivateSiswa(user_id, siswa_id, timestamp)
	if err != nil {
		log.Println(err)
		return errors.Wrap(err, "activate siswa by siswa Id on service")
	}

	return nil
}

func (siswa *siswaServiceimpl) DeactivateSiswa(siswa_id string, timestamp time.Time) error{
	if err := validations.ValidateUUID(siswa_id); err != nil {
		log.Println(err)
		return err
	}

	// check if user id found
	userfound, err := siswa.siswaRepository.CekSiswaExistByID(siswa_id)
	if err != nil && !userfound {
		log.Println(err)
		return err
	}

	// get user id
	user_id, err := siswa.siswaRepository.GetUserID(siswa_id)
	if err != nil {
		log.Println(err)
		return errors.Wrap(err, "getting user id by siswa id")
	}

	err = siswa.siswaRepository.DeactivateSiswa(user_id, siswa_id, timestamp)
	if err != nil {
		log.Println(err)
		return errors.Wrap(err, "deactivate siswa by siswa Id on service")
	}
	
	return nil
}