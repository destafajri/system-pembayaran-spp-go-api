package spp_service

import (
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

func (spp *sppServiceimpl) GetListSpp(meta *meta.Metadata) ([]model.GetListSppResponse, int, error) {
	resp, total, err := spp.sppRepository.GetListSppForAdmin(meta)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	return resp, total, nil
}