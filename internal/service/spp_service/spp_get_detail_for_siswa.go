package spp_service

import (
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/validations"
)

func (spp *sppServiceimpl) GetDetailSppForSiswa(spp_id, user_id string) (*model.GetDetailSppResponse, error) {
	if err := validations.ValidateUUID(spp_id); err != nil {
		return nil, err
	}

	// cek if spp exist
	ok, err := spp.sppRepository.CekSppExistByID(spp_id)
	if !ok && err != nil {
		log.Println(err)
		return nil, err
	}

	// get siswa id
	siswa_id, err := spp.sppRepository.GetSiswaIDByUserID(user_id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// cek if siswa id is match
	ok, err = spp.sppRepository.CekSppByIDAndSiswaIDIsMatch(spp_id, siswa_id)
	if !ok && err != nil {
		log.Println(err)
		return nil, err
	}

	resp, err := spp.sppRepository.GetDetailSpp(spp_id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return resp, nil
}