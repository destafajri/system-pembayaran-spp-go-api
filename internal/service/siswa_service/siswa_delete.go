package siswa_service

import (
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/validations"
	"github.com/pkg/errors"
)

func (siswa *siswaServiceimpl) DeleteSiswa(siswa_id string) error {
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

	err = siswa.siswaRepository.DeleteSiswa(user_id, siswa_id)
	if err != nil {
		log.Println(err)
		return errors.Wrap(err, "delete siswa on service")
	}

	return nil
}
