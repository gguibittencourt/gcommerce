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

func (o Writer) Create(ctx context.Context, order order.Order) error {
	tx := o.conn.Write.WithContext(ctx)
	return tx.Create(order).Error
}
