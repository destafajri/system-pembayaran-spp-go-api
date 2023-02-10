package spp_repository

import (
	"database/sql"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/repository"
)

type sppImplementation struct {
	db *sql.DB
}

func NewSppRepository(db *sql.DB) repository.SppRepository {
	return &sppImplementation{
		db: db,
	}
}

func (spp *sppImplementation) Sortable(field string) bool {
	switch field {
	case "created_at", "updated_at":
		return true
	default:
		return false
	}
}
