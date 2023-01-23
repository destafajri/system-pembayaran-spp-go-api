package entity

type SppEntity struct {
	ID         string `json:"id"`
	SiswaID    string `json:"siswa_id"`
	NoSpp      int    `json:"no_spp"`
	JatuhTempo string `json:"jatuh_tempo"`
	Jumlah     string `json:"jumlah"`
	IsActive   bool   `json:"is_active"`
	Timestamp
}
