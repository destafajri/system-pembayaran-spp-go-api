package siswa_service

import (
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/validations"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

func (siswa *siswaServiceimpl) GetListSiswaByKelas(kelas_id string, meta *meta.Metadata) ([]model.GetListSiswaByKelasResponse, int, error) {
	if err := validations.ValidateUUID(kelas_id); err != nil {
		log.Println(err)
		return nil, 0, err
	}

	resp, total, err := siswa.siswaRepository.GetListSiswaByKelas(kelas_id, meta)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	return resp, total, nil
}
