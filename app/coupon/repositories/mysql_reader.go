package repositories

import (
	"context"

	"github.com/gguibittencourt/gcommerce/app/coupon"
	"github.com/gguibittencourt/gcommerce/internal/database"
)

type Reader struct {
	conn database.Connection
}

func NewReader(conn database.Connection) Reader {
	return Reader{conn}
}

func (o Reader) FindByCode(ctx context.Context, code string) (coupon.Coupon, error) {
	tx := o.conn.Read.WithContext(ctx)
	model := couponModel{}
	err := tx.Where("code = ?", code).Find(&model).Error
	return model.toCoupon(), err
}
