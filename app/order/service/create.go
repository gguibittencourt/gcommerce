package service

import (
	"context"

	"github.com/gguibittencourt/gcommerce/app/order"
)

func CreateOrder(_ context.Context, order order.Order) error {
	if err := order.Validate(); err != nil {
		return err
	}
	return nil
}
