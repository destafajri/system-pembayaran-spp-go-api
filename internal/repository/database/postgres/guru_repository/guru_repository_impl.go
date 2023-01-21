package guru_repository

import (
	"database/sql"
	
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/repository"
)

type guruImplementation struct {
	db *sql.DB
}

func NewGuruRepository(db *sql.DB) repository.GuruRepository{
	return &guruImplementation{
		db: db,
	}
}

func (guru *guruImplementation) Sortable(field string) bool {
	switch field {
	case "created_at", "updated_at":
		return true
	default:
		return false
	}
}