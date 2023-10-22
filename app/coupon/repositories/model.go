package repositories

import (
	"time"

	"github.com/gguibittencourt/gcommerce/app/coupon"
	"github.com/gguibittencourt/gcommerce/internal/database"
)

type couponModel struct {
	database.Model
	Code           string
	ExpirationDate time.Time
	Percentage     float64
}

func (couponModel) TableName() string {
	return "coupon"
}

func (c couponModel) toCoupon() coupon.Coupon {
	return coupon.Coupon{
		CouponID:       c.ID,
		Code:           c.Code,
		Percentage:     c.Percentage,
		ExpirationDate: c.ExpirationDate,
	}
}
