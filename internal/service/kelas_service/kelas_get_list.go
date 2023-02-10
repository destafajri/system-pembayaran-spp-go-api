package kelas_service

import (
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

func (kelas *kelasServiceimpl) GetListKelas(meta *meta.Metadata) ([]model.GetListKelasResponse, int, error) {
	resp, total, err := kelas.kelasRepository.GetListKelas(meta)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	return resp, total, nil
}
