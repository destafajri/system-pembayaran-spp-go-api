package siswa_repository

import (
	"log"

	"github.com/pkg/errors"
)

func (siswa *siswaImplementation) GetKelasID(kelas string) (string, error) {
	var kelas_id string

	query := `SELECT id from kelas WHERE kelas = $1 LIMIT 1`

	rows := siswa.db.QueryRow(query, kelas)
	if err := rows.Scan(&kelas_id); err != nil {
		log.Println(err)
		return "", errors.New("kelas not found")
	}

	return kelas_id, nil
}