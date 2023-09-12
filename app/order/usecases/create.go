package usecases

import (
	"context"

	"github.com/gguibittencourt/gcommerce/app/freight"

	"github.com/gguibittencourt/gcommerce/app/order"
)

type (
	FreightService interface {
		Calculate(ctx context.Context, order order.Order) (freight.Freight, error)
	}
	CouponService interface {
		GetByCode(ctx context.Context, couponCode string) (order.Coupon, error)
	}
	Repository interface {
		Create(ctx context.Context, order order.Order) error
	}
	CreateUsecase struct {
		repository     Repository
		freightService FreightService
		couponService  CouponService
	}
)

func NewCreateUsecase(repository Repository, freightService FreightService, couponService CouponService) CreateUsecase {
	return CreateUsecase{
		repository:     repository,
		freightService: freightService,
		couponService:  couponService,
	}
}

func (c CreateUsecase) Create(ctx context.Context, order order.Order) error {
	coupon, err := c.couponService.GetByCode(ctx, order.Coupon.Code)
	if err != nil {
		return err
	}
	order.Coupon = coupon
	if err := order.Validate(); err != nil {
		return err
	}
	f, err := c.freightService.Calculate(ctx, order)
	if err != nil {
		return err
	}
	order.Freight = f
	return c.repository.Create(ctx, order)
}

//TODO fill the coupon
//TODO validate the order
//TODO fill the freight
//TODO fill the code
//TODO fill the total
//TODO fill the status
