package siswa_service

import (
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

func (siswa *siswaServiceimpl) GetListSiswa(role string, meta *meta.Metadata) ([]model.GetListSiswaResponse, int, error) {
	if role == "admin" {
		resp, total, err := siswa.siswaRepository.GetListSiswaForAdmin(meta)
		if err != nil {
			log.Println(err)
			return nil, 0, err
		}
	
		return resp, total, nil
	} else {
		resp, total, err := siswa.siswaRepository.GetListSiswaNonAdmin(meta)
		if err != nil {
			log.Println(err)
			return nil, 0, err
		}
	
		return resp, total, nil
	}
}
