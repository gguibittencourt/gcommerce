package entity

import (
	"time"
)

type Coupon struct {
	Code       string
	Percentage float64
	ExpireDate time.Time
}

func (c Coupon) ApplyDiscount(total float64, date time.Time) float64 {
	if c.ExpireDate.After(date) {
		return total
	}
	return total - ((c.Percentage * total) / 100)
}
