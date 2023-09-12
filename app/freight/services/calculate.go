package services

import (
	"context"

	"github.com/gguibittencourt/gcommerce/app/freight"
	"github.com/gguibittencourt/gcommerce/app/order"
)

type (
	Service struct {
		repository Repository
	}

	Repository interface {
		Calculate(ctx context.Context, order order.Order) (freight.Freight, error)
	}
)

func NewCalculateService(repository Repository) Service {
	return Service{
		repository: repository,
	}
}

func (s Service) Calculate(ctx context.Context, order order.Order) (freight.Freight, error) {
	return s.repository.Calculate(ctx, order)
}
