package kelas_service

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/repository"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/service"
)

type kelasServiceimpl struct {
	kelasRepository repository.KelasRepository
}

func NewkelasService(kelasRepository *repository.KelasRepository) service.KelasService {
	return &kelasServiceimpl{
		kelasRepository: *kelasRepository,
	}
}