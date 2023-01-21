package entity

type SiswaEntity struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	KelasID  string `json:"kelas_id"`
	NIS      string `json:"nis"`
	Nama     string `json:"nama"`
	Angkatan string `json:"angkatan"`
	IsActive bool   `json:"is_active"`
}
