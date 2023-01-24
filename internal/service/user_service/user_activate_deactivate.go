package user_service

import (
	"log"
	"time"

	validations "github.com/destafajri/system-pembayaran-spp-go-api/internal/validations"
	"github.com/pkg/errors"
)

func (user *userServiceimpl) ActivateUser(id string, timestamp time.Time) error {
	if err := validations.ValidateUUID(id); err != nil {
		log.Println(err)
		return err
	}

	// check if user id found
	userfound, err := user.userRepository.CekUserExistByID(id)
	if err != nil && !userfound {
		log.Println(err)
		return err
	}

	// get user role
	role, err := user.userRepository.GetRoleInformation(id)
	if err != nil {
		log.Println(err)
		return err
	}

	if role == "guru" {
		guru_id, err := user.userRepository.GetGuruID(id)
		if err != nil {
			log.Println(err)
			return err
		}

		err = user.userRepository.ActivateGuru(id, guru_id, timestamp)
		if err != nil {
			log.Println(err)
			return errors.Wrap(err, "activate guru by user Id on service")
		}

		return nil
	} else if role == "siswa" {
		siswa_id, err := user.userRepository.GetSiswaID(id)
		if err != nil {
			log.Println(err)
			return err
		}

		err = user.userRepository.ActivateSiswa(id, siswa_id, timestamp)
		if err != nil {
			log.Println(err)
			return errors.Wrap(err, "activate siswa by user Id on service")
		}

		return nil
	} else {
		err = user.userRepository.ActivateUser(id, timestamp)
		if err != nil {
			log.Println(err)
			return errors.Wrap(err, "activate user by user Id on service")
		}

		return nil
	}
}

func (user *userServiceimpl) DeactivateUser(id string, timestamp time.Time) error {
	if err := validations.ValidateUUID(id); err != nil {
		log.Println(err)
		return err
	}

	// check if user id found
	userfound, err := user.userRepository.CekUserExistByID(id)
	if err != nil && !userfound {
		log.Println(err)
		return err
	}

	// get user role
	role, err := user.userRepository.GetRoleInformation(id)
	if err != nil {
		log.Println(err)
		return err
	}

	if role == "guru" {
		guru_id, err := user.userRepository.GetGuruID(id)
		if err != nil {
			log.Println(err)
			return err
		}

		err = user.userRepository.DeactivateGuru(id, guru_id, timestamp)
		if err != nil {
			log.Println(err)
			return errors.Wrap(err, "deactivate guru by user Id on service")
		}

		return nil
	} else if role == "siswa" {
		siswa_id, err := user.userRepository.GetSiswaID(id)
		if err != nil {
			log.Println(err)
			return err
		}

		err = user.userRepository.DeactivateSiswa(id, siswa_id, timestamp)
		if err != nil {
			log.Println(err)
			return errors.Wrap(err, "deactivate siswa by user Id on service")
		}

		return nil
	} else {
		err = user.userRepository.DeactivateUser(id, timestamp)
		if err != nil {
			log.Println(err)
			return errors.Wrap(err, "deactivate user by user Id on service")
		}

		return nil
	}
}
