package siswa_service

import (
	"log"
	"strings"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/validations"
	"github.com/pkg/errors"
)

func (siswa *siswaServiceimpl) UpdateKelasSiswa(siswa_id string, input *model.UpdateKelasSiswaRequest) (*model.UpdateKelasSiswaResponse, error) {
	if err := validations.ValidateUUID(siswa_id); err != nil {
		log.Println(err)
		return nil, err
	}

	// validation input
	err := model.ValidateUpdateKelasSiswaInput(input)
	if err != nil {
		return nil, err
	}

	angkatan, err := validations.YearValidation(input.Angkatan)
	if err != nil && !angkatan {
		return nil, err
	}

	// check if siswa id found
	userfound, err := siswa.siswaRepository.CekSiswaExistByID(siswa_id)
	if err != nil && !userfound {
		log.Println(err)
		return nil, err
	}

	// get kelas id
	kelas_id, err := siswa.siswaRepository.GetKelasID(strings.ToUpper(input.Kelas))
	if err != nil {
		log.Println(err)
		return nil, errors.Wrap(err, "getting user id by siswa id")
	}

	err = siswa.siswaRepository.UpdateKelasSiswa(siswa_id, kelas_id, input.Angkatan)
	if err != nil {
		log.Println(err)
		return nil, errors.Wrap(err, "update kelas siswa on service")
	}

	siswaDetail, err := siswa.siswaRepository.GetDetailSiswaForAdmin(siswa_id)
	if err != nil {
		log.Println(err)
		return nil, errors.New("getting detail siswa on service update kelas")
	}

	resp := model.UpdateKelasSiswaResponse{
		ID:       siswaDetail.ID,
		Nama:     siswaDetail.Nama,
		NIS:      siswaDetail.NIS,
		Kelas:    siswaDetail.Kelas,
		Angkatan: siswaDetail.Angkatan,
		IsActive: siswaDetail.IsActive,
	}

	return &resp, nil
}
