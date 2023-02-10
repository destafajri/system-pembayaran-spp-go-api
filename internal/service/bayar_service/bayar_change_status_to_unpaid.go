package bayar_service

import (
	"errors"
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"
)

func (bayar *bayarServiceimpl) UnpaidSpp(input *model.BayarSppRequest) (*model.BayarSppResponse, error) {
	// validation input
	err := model.ValidateCallbackBayarInput(input)
	if err != nil {
		log.Println()
		return nil, err
	}

	// get spp value
	spp, err := bayar.bayarRepository.GetSppInfoBySppNumber(input.NoSpp)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// cek if status unpaid
	status, _ := bayar.bayarRepository.GetStatusInfo(spp.ID)
	if status != "paid" {
		return nil, errors.New("siswa haven't paid that spp")
	}

	err = bayar.bayarRepository.UnpaidSpp(spp.ID)
	if err != nil {
		return nil, err
	}

	resp := model.BayarSppResponse{
		ID:         spp.ID,
		SiswaID:    spp.SiswaID,
		NoSpp:      spp.NoSpp,
		JatuhTempo: spp.JatuhTempo,
		Jumlah:     spp.Jumlah,
		Status:     "unpaid",
	}

	return &resp, nil
}
