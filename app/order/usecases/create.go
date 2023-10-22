package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/gguibittencourt/gcommerce/app/coupon"
	"github.com/gguibittencourt/gcommerce/app/freight"

	"github.com/gguibittencourt/gcommerce/app/order"
)

type (
	FreightReader interface {
		Calculate(ctx context.Context, order order.Order) (freight.Freight, error)
	}
	CouponReader interface {
		FindByCode(ctx context.Context, code string) (coupon.Coupon, error)
	}
	Writer interface {
		Save(ctx context.Context, order order.Order) error
	}
	Publisher interface {
		Publish(ctx context.Context, msg any) error
	}
	CreateUsecase struct {
		writer         Writer
		freightService FreightReader
		couponReader   CouponReader
		publisher      Publisher
	}
)

func NewCreateUsecase(writer Writer, freightReader FreightReader, couponReader CouponReader, publisher Publisher) CreateUsecase {
	return CreateUsecase{writer, freightReader, couponReader, publisher}
}

func (u CreateUsecase) Execute(ctx context.Context, o order.Order) (err error) {
	o.Coupon, err = u.couponReader.FindByCode(ctx, o.Coupon.Code)
	if err != nil {
		return err
	}
	if err = o.Validate(); err != nil {
		return err
	}
	o.Freight, err = u.freightService.Calculate(ctx, o)
	if err != nil {
		return err
	}
	o.Status = order.StatusPending
	o.Code = fmt.Sprintf("%d", time.Now().UnixNano())
	if err = u.writer.Save(ctx, o); err != nil {
		return err
	}
	return u.publisher.Publish(ctx, o)
}
