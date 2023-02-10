package guru_service

import (
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

func (guru *guruServiceimpl) GetListGuru(role string, meta *meta.Metadata) ([]model.GetListGuruResponse, int, error) {
	if role == "admin" {
		resp, total, err := guru.guruRepository.GetListGuruAdmin(meta)
		if err != nil {
			log.Println(err)
			return nil, 0, err
		}
	
		return resp, total, nil
	} else {
		resp, total, err := guru.guruRepository.GetListGuruNonAdmin(meta)
		if err != nil {
			log.Println(err)
			return nil, 0, err
		}
	
		return resp, total, nil
	}
}
