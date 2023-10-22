package usecases_test

//func TestCreateOrder(t *testing.T) {
//	tests := []struct {
//		name     string
//		order    entity.Order
//		expected error
//	}{
//		{
//			name: "given a order without discount and three items, should return nil",
//			order: entity.Order{
//				Items: entity.Items{
//					buildItem(1, 10),
//					buildItem(2, 20),
//					buildItem(3, 30),
//				},
//			},
//			expected: nil,
//		},
//		{
//			name:     "given a order without items, should return error",
//			order:    entity.Order{},
//			expected: errors.New("order without items"),
//		},
//		{
//			name: "given a order with total discount, should return error",
//			order: entity.Order{
//				Items: entity.Items{
//					buildItem(1, 10),
//					buildItem(2, 20),
//					buildItem(3, 30),
//				},
//				Discount: 1,
//			},
//			expected: errors.New("the total order price is invalid"),
//		},
//	}
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			err := writer.CreateOrder(context.TODO(), test.order)
//			require.Equal(t, test.expected, err)
//		})
//	}
//}
