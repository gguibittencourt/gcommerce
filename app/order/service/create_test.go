package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gguibittencourt/gcommerce/app/order"
	"github.com/gguibittencourt/gcommerce/app/order/service"
)

func TestCreateOrder(t *testing.T) {
	tests := []struct {
		name     string
		order    order.Order
		expected error
	}{
		{
			name: "given a order without discount and three items, should return nil",
			order: order.Order{
				Items: order.Items{
					buildItem(1, 10),
					buildItem(2, 20),
					buildItem(3, 30),
				},
			},
			expected: nil,
		},
		{
			name:     "given a order without items, should return error",
			order:    order.Order{},
			expected: errors.New("order without items"),
		},
		{
			name: "given a order with total discount, should return error",
			order: order.Order{
				Items: order.Items{
					buildItem(1, 10),
					buildItem(2, 20),
					buildItem(3, 30),
				},
				Discount: 1,
			},
			expected: errors.New("the total order price is invalid"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := service.CreateOrder(context.TODO(), test.order)
			require.Equal(t, test.expected, err)
		})
	}
}

func buildItem(amount uint32, price float64) order.Item {
	return order.Item{
		Amount:      amount,
		Description: "Item",
		Price:       price,
	}
}
