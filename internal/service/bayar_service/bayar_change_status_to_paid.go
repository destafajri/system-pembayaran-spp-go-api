package bayar_service

import (
	"errors"
	"log"
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"
	"github.com/google/uuid"
)

func (bayar *bayarServiceimpl) PaidSpp(input *model.BayarSppRequest) (*model.BayarSppResponse, error) {
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
	if status != "unpaid" {
		return nil, errors.New("siswa already paid that spp")
	}

	payload := entity.BayarEntity{
		ID:           uuid.New().String(),
		SppID:        spp.ID,
		TanggalBayar: time.Now().Local().Format("2006-01-02"),
	}

	err = bayar.bayarRepository.PaidSpp(&payload)
	if err != nil {
		return nil, err
	}

	resp := model.BayarSppResponse{
		ID:         spp.ID,
		SiswaID:    spp.SiswaID,
		NoSpp:      spp.NoSpp,
		JatuhTempo: spp.JatuhTempo,
		Jumlah:     spp.Jumlah,
		Status:     "paid",
	}

	return &resp, nil
}
