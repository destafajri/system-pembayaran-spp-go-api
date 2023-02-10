package siswa_repository

import (
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/pkg/errors"
)

func (siswa *siswaImplementation) CekNIS(nim int) (string, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	var siswa_id string

	query := `SELECT id FROM siswa WHERE nis = $1 LIMIT 1;`

	rows := siswa.db.QueryRow(query, nim)
	if err := rows.Scan(&siswa_id); err != nil {
		log.Println(err)
		return "", errors.New("siswa not found")
	}

	return siswa_id, nil
}