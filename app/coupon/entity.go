package coupon

import "time"

type Coupon struct {
	CouponID       uint64    `json:"coupon_id"`
	Code           string    `json:"code"`
	Percentage     float64   `json:"percentage"`
	ExpirationDate time.Time `json:"expiration_date"`
}

func (c Coupon) ApplyDiscount(total float64, date time.Time) float64 {
	if c.ExpirationDate.After(date) {
		return total
	}
	return total - ((c.Percentage * total) / 100)
}
