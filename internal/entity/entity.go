package entity

import "time"

type Error string

const (
	ErrPermissionNotAllowed = Error("permission.not_allowed")
)

func (e Error) Error() string {
	return string(e)
}

type Timestamp struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
