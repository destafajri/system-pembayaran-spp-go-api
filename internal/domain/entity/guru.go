package entity

type GuruEntity struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Nama     string `json:"nama"`
	IsActive bool   `json:"is_active"`
}
