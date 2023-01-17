package user_repository

import (
	"database/sql"
	
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/repository"
)

type userImplementation struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository{
	return &userImplementation{
		db: db,
	}
}