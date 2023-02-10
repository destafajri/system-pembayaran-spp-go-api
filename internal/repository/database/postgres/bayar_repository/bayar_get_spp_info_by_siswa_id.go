package bayar_repository

import (
	"errors"
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/config"
	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/entity"
)

func (bayar *bayarImplementation) GetSppInfoBySppNumber(no_spp string) (*entity.SppEntity, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	var spp entity.SppEntity

	query := `SELECT * FROM spp WHERE no_spp = $1 AND is_active is true LIMIT 1`

	rows := bayar.db.QueryRow(query, no_spp)
	if err := rows.Scan(
			&spp.ID,
			&spp.SiswaID,
			&spp.NoSpp,
			&spp.JatuhTempo,
			&spp.Jumlah,
			&spp.IsActive,
			&spp.CreatedAt,
			&spp.UpdatedAt,
			&spp.DeletedAt,
		); err != nil {
		log.Println(err)
		return nil, errors.New("spp not found")
	}

	return &spp, nil
}