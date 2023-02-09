package service

import "github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"

type BayarService interface {
	PaidSpp(input *model.BayarSppRequest) (*model.BayarSppResponse, error)
}