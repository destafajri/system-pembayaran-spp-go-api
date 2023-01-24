package repository

import (
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/entity"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/model"
	"github.com/destafajri/system-pembayaran-spp-go-api/meta"
)

type UserRepository interface {
	Login(username string) (*entity.UserEntity, error)
	CreateAdmin(users *entity.UserEntity) (*model.CreateAdminResponse, error)

	GetListUser(meta *meta.Metadata) ([]model.GetListUserResponse, int, error)
	GetDetailUser(id string) (*model.GetDetailUser, error)

	ActivateUser(id string, timestamp time.Time) error
	DeactivateUser(id string, timestamp time.Time) error
	ActivateGuru(user_id, guru_id string, timestamp time.Time) error
	DeactivateGuru(user_id, guru_id string, timestamp time.Time) error
	ActivateSiswa(user_id, siswa_id string, timestamp time.Time) error
	DeactivateSiswa(user_id, siswa_id string, timestamp time.Time) error
	
	CekUserExistByID(id string) (bool, error)
	GetRoleInformation(id string) (string, error)
	GetGuruID(user_id string) (string, error)
	GetSiswaID(user_id string) (string, error)
}
