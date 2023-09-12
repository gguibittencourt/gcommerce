package freight

import "time"

type Freight struct {
	FreightID     uint64        `json:"freight_id"`
	Price         float64       `json:"price"`
	DurationInMin time.Duration `json:"duration_in_min"`
	ETA           time.Time     `json:"eta"`
}
