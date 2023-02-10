package bayar_repository

import (
	"database/sql"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/repository"
)

type bayarImplementation struct {
	db *sql.DB
}

func NewBayarRepository(db *sql.DB) repository.BayarRepository {
	return &bayarImplementation{
		db: db,
	}
}

func (bayar *bayarImplementation) Sortable(field string) bool {
	switch field {
	case "created_at", "updated_at":
		return true
	default:
		return false
	}
}