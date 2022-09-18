package order

import "errors"

type (
	Order struct {
		OrderID  uint64
		CPF      string
		Discount float64
		Items    Items
	}
	Item struct {
		Amount      uint32
		Description string
		Price       float64
	}
	Items []Item
)

func (o Order) Total() float64 {
	total := float64(0)
	for i := range o.Items {
		item := o.Items[i]
		total += item.Price * float64(item.Amount)
	}
	if o.Discount > 0 {
		total -= total * o.Discount
	}
	return total
}

func (o Order) Validate() error {
	if len(o.Items) == 0 {
		return errors.New("order without items")
	}
	total := o.Total()
	if total <= 0 {
		return errors.New("the total order price is invalid")
	}
	return nil
}
