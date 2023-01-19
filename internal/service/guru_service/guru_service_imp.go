package guru_service

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/repository"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/service"
)

type guruServiceimpl struct {
	guruRepository repository.GuruRepository
}

func NewUserService(guruRepository *repository.GuruRepository) service.GuruService {
	return &guruServiceimpl{
		guruRepository: *guruRepository,
	}
}