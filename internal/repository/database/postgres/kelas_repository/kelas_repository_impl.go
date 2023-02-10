package kelas_repository

import (
	"database/sql"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/repository"
)

type kelasImplementation struct {
	db *sql.DB
}

func NewkelasRepository(db *sql.DB) repository.KelasRepository {
	return &kelasImplementation{
		db: db,
	}
}

func (kelas *kelasImplementation) Sortable(field string) bool {
	switch field {
	case "created_at", "updated_at":
		return true
	default:
		return false
	}
}
