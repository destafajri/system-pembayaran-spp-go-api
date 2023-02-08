package spp_service

import (
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

func (spp *sppServiceimpl) GetListSpp(kelasparam string, meta *meta.Metadata) ([]model.GetListSppResponse, int, error) {
	resp, total, err := spp.sppRepository.GetListSppForAdmin(kelasparam, meta)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	return resp, total, nil
}