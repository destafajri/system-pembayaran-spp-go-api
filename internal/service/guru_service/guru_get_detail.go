package guru_service

import (
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/validations"
	"github.com/pkg/errors"
)

func (guru *guruServiceimpl) GetDetailGuru(role, guru_id string) (*model.GetDetailGuruResponse, error) {
	if err := validations.ValidateUUID(guru_id); err != nil {
		log.Println(err)
		return nil, err
	}

	// check if guru id found
	gurufound, err := guru.guruRepository.CekGuruExistByID(guru_id)
	if err != nil && !gurufound {
		log.Println(err)
		return nil, err
	}

	if role == "admin" {
		guruDetail, err := guru.guruRepository.GetDetailGuruAdmin(guru_id)
		if err != nil {
			log.Println(err)
			return nil, errors.Wrap(err, "find guru by guru Id on service")
		}
	
		return guruDetail, nil
	} else {
		guruDetail, err := guru.guruRepository.GetDetailGuruNonAdmin(guru_id)
		if err != nil {
			log.Println(err)
			return nil, errors.Wrap(err, "find guru by guru Id on service")
		}
	
		return guruDetail, nil
	}
}
