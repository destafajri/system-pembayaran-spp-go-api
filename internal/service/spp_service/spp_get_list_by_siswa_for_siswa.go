package spp_service

import (
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

func (spp *sppServiceimpl) GetListSppBySiswaForSiswa(user_id string, meta *meta.Metadata) ([]model.GetListSppResponse, int, error) {
	siswa_id, err := spp.sppRepository.GetSiswaIDByUserID(user_id)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	resp, total, err := spp.sppRepository.GetListSppBySiswa(siswa_id, meta)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	return resp, total, nil
}
