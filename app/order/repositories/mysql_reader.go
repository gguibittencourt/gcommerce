package repositories

import (
	"context"

	"github.com/gguibittencourt/gcommerce/app/order"
	"github.com/gguibittencourt/gcommerce/internal/database"
)

type Reader struct {
	conn database.Connection
}

func NewReader(conn database.Connection) Reader {
	return Reader{conn}
}

func (o Reader) FindByID(ctx context.Context, id uint64) (order.Order, error) {
	tx := o.conn.Read.WithContext(ctx)
	model := orderModel{}
	err := tx.Take(&model, id).Error
	return model.toOrder(), err
}
