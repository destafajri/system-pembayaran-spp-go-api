package siswa_repository

import (
	"errors"
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
)

func (siswa *siswaImplementation) CreateSiswa(userReq *entity.UserEntity, siswaReq *entity.SiswaEntity) (*model.CreateSiswaResponse, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

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

	_, err := siswa.db.Exec(query1, values...)
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

	_, err = siswa.db.Exec(query2, value...)
	if err != nil {
		log.Println(err)
		return nil, errors.New("guru already exist")
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
