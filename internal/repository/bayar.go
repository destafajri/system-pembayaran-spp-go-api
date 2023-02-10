package repository

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/entity"
)

type BayarRepository interface {
	PaidSpp(input *entity.BayarEntity) error
	UnpaidSpp(spp_id string) error

	GetSppInfoBySppNumber(no_spp string) (*entity.SppEntity, error)
	GetStatusInfo(spp_id string) (string, error)
}
