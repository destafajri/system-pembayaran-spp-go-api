package siswa_repository

import (
	"database/sql"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/repository"
)

type siswaImplementation struct {
	db *sql.DB
}

func NewSiswaRepository(db *sql.DB) repository.SiswaRepository{
	return &siswaImplementation{
		db: db,
	}
}

func (siswa *siswaImplementation) Sortable(field string) bool {
	switch field {
	case "created_at", "updated_at":
		return true
	default:
		return false
	}
}