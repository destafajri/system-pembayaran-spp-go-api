package siswa_service

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/repository"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/service"
)

type siswaServiceimpl struct {
	siswaRepository repository.SiswaRepository
}

func NewSiswaService(siswaRepository *repository.SiswaRepository) service.SiswaService {
	return &siswaServiceimpl{
		siswaRepository: *siswaRepository,
	}
}