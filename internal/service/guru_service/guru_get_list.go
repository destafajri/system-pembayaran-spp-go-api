package guru_service

import (
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

func (guru *guruServiceimpl) GetListGuru(meta *meta.Metadata) ([]model.GetListGuruResponse, int, error) {
	resp, total, err := guru.guruRepository.GetListGuru(meta)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	return resp, total, nil
}
