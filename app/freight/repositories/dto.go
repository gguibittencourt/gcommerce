package repositories

import (
	"time"

	"github.com/gguibittencourt/gcommerce/app/freight"
)

type freightDTO struct {
	Code          string        `json:"code"`
	FreightID     uint64        `json:"freight_id"`
	Price         float64       `json:"price"`
	DurationInMin time.Duration `json:"duration_in_min"`
	ETA           time.Time     `json:"eta"`
}

func (f freightDTO) toEntity() freight.Freight {
	return freight.Freight{
		Code:          f.Code,
		FreightID:     f.FreightID,
		Price:         f.Price,
		DurationInMin: f.DurationInMin,
		ETA:           f.ETA,
	}
}
