package order_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/gguibittencourt/gcommerce/app/order"
)

func TestCreateOrder(t *testing.T) {
	var (
		validCPF           = "470.508.590-60"
		invalidCPF         = "111.111.111-11"
		normalDate         = time.Date(2022, 9, 25, 0, 0, 0, 0, time.UTC)
		expiredDate        = time.Date(2022, 9, 27, 0, 0, 0, 0, time.UTC)
		allDiscount        = buildCoupon(100, normalDate)
		allDiscountExpired = buildCoupon(100, expiredDate)
		halfDiscount       = buildCoupon(50, normalDate)
		threeOrderItems    = []order.Item{
			buildOrderItem(1, 1, 10),
			buildOrderItem(2, 2, 20),
			buildOrderItem(3, 3, 30),
		}
	)
	tests := []struct {
		name     string
		order    order.Order
		expected error
	}{
		{
			name: "given an order with discount and three items, should return nil",
			order: order.Order{
				CPF:    validCPF,
				Items:  threeOrderItems,
				Coupon: halfDiscount,
			},
			expected: nil,
		},
		{
			name: "given an order with expired coupon with total discount, should return error",
			order: order.Order{
				CPF:    validCPF,
				Items:  threeOrderItems,
				Coupon: allDiscountExpired,
			},
			expected: nil,
		},
		{
			name: "given an order with invalid CPF, should return error",
			order: order.Order{
				CPF: invalidCPF,
			},
			expected: errors.New("invalid CPF"),
		},
		{
			name: "given an order without items, should return error",
			order: order.Order{
				CPF: validCPF,
			},
			expected: errors.New("order without items"),
		},
		{
			name: "given an order with total discount, should return error",
			order: order.Order{
				CPF:    validCPF,
				Items:  threeOrderItems,
				Coupon: allDiscount,
			},
			expected: errors.New("the total order price is invalid"),
		},
		{
			name: "given an order with invalid quantity item, should return error",
			order: order.Order{
				CPF: validCPF,
				Items: order.Items{
					buildOrderItem(1, 0, 10),
				},
				Coupon: allDiscount,
			},
			expected: errors.New("invalid quantity of 1"),
		},
		{
			name: "given an order with duplicated item, should return error",
			order: order.Order{
				CPF: validCPF,
				Items: order.Items{
					buildOrderItem(1, 1, 10),
					buildOrderItem(1, 1, 10),
				},
			},
			expected: errors.New("duplicated item 1"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.order.Validate()
			require.Equal(t, test.expected, err)
		})
	}
}

func buildOrderItem(itemID uint64, amount uint32, price float64) order.Item {
	return order.Item{
		ItemID: itemID,
		Amount: amount,
		Price:  price,
	}
}

func buildCoupon(percentage float64, date time.Time) order.Coupon {
	return order.Coupon{
		Code:           "code",
		Percentage:     percentage,
		ExpirationDate: date,
	}
}
