package spp_service

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/repository"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/service"
)

type sppServiceimpl struct {
	sppRepository repository.SppRepository
}

func NewSppService(sppRepository *repository.SppRepository) service.SppService {
	return &sppServiceimpl{
		sppRepository: *sppRepository,
	}
}