package repositories

import (
	"context"

	"github.com/gguibittencourt/gcommerce/app/order"
	"github.com/gguibittencourt/gcommerce/internal/database"
)

type Writer struct {
	conn database.Connection
}

func NewWriter(conn database.Connection) Writer {
	return Writer{conn}
}

func (o Writer) Save(ctx context.Context, order order.Order) error {
	model := toOrderModel(order)
	tx := o.conn.Write.WithContext(ctx)
	return tx.Save(&model).Error
}
