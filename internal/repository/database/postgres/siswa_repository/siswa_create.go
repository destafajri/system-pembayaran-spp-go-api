package siswa_repository

import (
	"fmt"
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"
	"github.com/pkg/errors"
)

func (siswa *siswaImplementation) CreateSiswa(userReq *entity.UserEntity, siswaReq *entity.SiswaEntity) (*model.CreateSiswaResponse, error) {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()

	// Create a helper function for preparing failure results.
	fail := func(err error) error {
		return fmt.Errorf("CREATE SISWA: %v", err)
	}

	// Get a Tx for making transaction requests.
	tx, err := siswa.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fail(err)
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	query1 := `INSERT INTO users(
			id,
			email,
			username,
			password,
			role,
			is_active,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	values := []interface{}{
		userReq.ID,
		userReq.Email,
		userReq.Username,
		userReq.Password,
		userReq.Role,
		userReq.IsActive,
		userReq.CreatedAt,
		userReq.UpdatedAt,
	}

	_, err = siswa.db.Exec(query1, values...)
	if err != nil {
		log.Println(err)
		return nil, errors.New("user already exist")
	}

	query2 := `INSERT INTO siswa(
			id,
			user_id,
			kelas_id,
			nis,
			nama,
			angkatan,
			is_active
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	value := []interface{}{
		siswaReq.ID,
		siswaReq.UserID,
		siswaReq.KelasID,
		siswaReq.NIS,
		siswaReq.Nama,
		siswaReq.Angkatan,
		siswaReq.IsActive,
	}

	_, err = siswa.db.ExecContext(ctx, query2, value...)
	if err != nil {
		log.Println(err)
		return nil, errors.New("siswa already exist")
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return nil, fail(err)
	}

	resp := model.CreateSiswaResponse{
		ID:        siswaReq.ID,
		UserID:    siswaReq.UserID,
		KelasID:   siswaReq.KelasID,
		Email:     userReq.Email,
		Username:  userReq.Username,
		Role:      userReq.Role,
		NIS:       siswaReq.NIS,
		Nama:      siswaReq.Nama,
		Angkatan:  siswaReq.Angkatan,
		IsActive:  siswaReq.IsActive,
		Timestamp: userReq.Timestamp,
	}

	return &resp, nil
}
