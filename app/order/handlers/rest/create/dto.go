package create

import (
	"github.com/gguibittencourt/gcommerce/app/order"
)

type (
	Payload struct {
		CPF    string `json:"cpf"`
		Items  []Item `json:"items"`
		Coupon string `json:"coupon"`
	}

	Item struct {
		ProductID uint64
		Amount    uint32
		Price     float64
	}
)

func (p Payload) toOrder() order.Order {
	items := make(order.Items, len(p.Items))
	for i := range p.Items {
		items[i] = p.Items[i].toItem()
	}
	return order.Order{
		CPF:   p.CPF,
		Items: items,
		Coupon: order.Coupon{
			Code: p.Coupon,
		},
	}
}

func (i Item) toItem() order.Item {
	return order.Item{
		Amount:    i.Amount,
		Price:     i.Price,
		ProductID: i.ProductID,
	}
}
