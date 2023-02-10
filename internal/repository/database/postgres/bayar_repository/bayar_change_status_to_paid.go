package bayar_repository

import (
	"errors"
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/entity"
)

func (bayar *bayarImplementation) PaidSpp(input *entity.BayarEntity) error {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
			INSERT INTO bayar(
				id,
				spp_id,
				tanggal_bayar
			)
			VALUES ($1, $2, $3)
		`
	values := []interface{}{
		input.ID,
		input.SppID,
		input.TanggalBayar,
	}

	_, err := bayar.db.Exec(query, values...)
	if err != nil {
		log.Println(err)
		return errors.New("spp already exist")
	}

	return nil
}