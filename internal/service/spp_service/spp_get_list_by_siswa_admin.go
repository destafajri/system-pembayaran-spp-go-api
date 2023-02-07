package spp_service

import (
	"errors"
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

func(spp *sppServiceimpl) GetListSppBySiswaForAdmin(siswa_id string, meta *meta.Metadata) ([]model.GetListSppResponse, int, error) {
	if siswa_id == "" {
		return nil, 0, errors.New("siswa id cannot be blank")
	}

	resp, total, err := spp.sppRepository.GetListSppBySiswa(siswa_id, meta)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	return resp, total, nil
}