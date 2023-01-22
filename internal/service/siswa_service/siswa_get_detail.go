package siswa_service

import (
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/validations"
	"github.com/pkg/errors"
)

func (siswa *siswaServiceimpl) GetDetailSiswa(role, siswa_id string) (*model.GetDetailSiswaResponse, error) {
	if err := validations.ValidateUUID(siswa_id); err != nil {
		log.Println(err)
		return nil, err
	}

	// check if siswa id found
	siswafound, err := siswa.siswaRepository.CekSiswaExistByID(siswa_id)
	if err != nil && !siswafound {
		log.Println(err)
		return nil, err
	}

	if role == "admin" {
		siswaDetail, err := siswa.siswaRepository.GetDetailSiswaForAdmin(siswa_id)
		if err != nil {
			log.Println(err)
			return nil, errors.Wrap(err, "find siswa by siswa Id on service")
		}
	
		return siswaDetail, nil
	} else {
		siswaDetail, err := siswa.siswaRepository.GetDetailSiswaNonAdmin(siswa_id)
		if err != nil {
			log.Println(err)
			return nil, errors.Wrap(err, "find siswa by siswa Id on service")
		}
	
		return siswaDetail, nil
	}
}
