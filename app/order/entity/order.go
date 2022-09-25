package entity

import (
	"errors"
	"time"

	"github.com/gguibittencourt/gcommerce/pkg/validator"
)

type (
	Order struct {
		OrderID uint64
		CPF     string
		Coupon  Coupon
		Items   OrderItems
		Date    time.Time
	}
)

func (o Order) Validate() error {
	if !validator.IsValidCPF(o.CPF) {
		return errors.New("invalid CPF")
	}
	if err := o.Items.Validate(); err != nil {
		return err
	}
	total := o.Total()
	if total <= 0 {
		return errors.New("the total order price is invalid")
	}
	return nil
}

func (o Order) Total() float64 {
	total := o.Items.Total()
	if o.Coupon.Code != "" {
		total = o.Coupon.ApplyDiscount(total, o.Date)
	}
	return total
}
