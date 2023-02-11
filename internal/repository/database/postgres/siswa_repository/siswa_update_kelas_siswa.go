package siswa_repository

import (
	"errors"
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
)

func (siswa *siswaImplementation) UpdateKelasSiswa(siswa_id, kelas_id, angkatan string) error {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
	UPDATE siswa SET
		kelas_id = $1,
		angkatan = $2
	WHERE 
		id = $3
	`
	value := []interface{}{
		kelas_id,
		angkatan,
		siswa_id,
	}

	_, err := siswa.db.Exec(query, value...)
	if err != nil {
		log.Println(err)
		return errors.New("error update kelas siswa")
	}

	return nil
}
