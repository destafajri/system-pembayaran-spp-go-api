package bayar_service

import (
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/repository"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/service"
)

type bayarServiceimpl struct {
	bayarRepository repository.BayarRepository
}

func NewBayarService(bayarRepository *repository.BayarRepository) service.BayarService {
	return &bayarServiceimpl{
		bayarRepository: *bayarRepository,
	}
}