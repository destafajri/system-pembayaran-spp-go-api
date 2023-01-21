package entity

type KelasEntity struct {
	ID       string `json:"id"`
	GuruID   string `json:"guru_id"`
	Kelas    string `json:"kelas"`
	Timestamp
}