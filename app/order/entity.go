package order

import (
	"errors"
	"fmt"
	"time"

	"github.com/gguibittencourt/gcommerce/app/coupon"
	"github.com/gguibittencourt/gcommerce/app/freight"

	"github.com/gguibittencourt/gcommerce/pkg/validator"
)

const (
	StatusPending = "pending"
)

type (
	Order struct {
		OrderID   uint64          `json:"order_id"`
		Code      string          `json:"code"`
		CPF       string          `json:"cpf"`
		Status    string          `json:"status"`
		Items     Items           `json:"items"`
		Coupon    coupon.Coupon   `json:"coupon"`
		Freight   freight.Freight `json:"freight"`
		CreatedAt time.Time       `json:"created_at"`
		UpdatedAt time.Time       `json:"updated_at"`
	}

	Item struct {
		ItemID    uint64  `json:"item_id"`
		ProductID uint64  `json:"product_id"`
		Amount    uint32  `json:"amount"`
		Price     float64 `json:"price"`
	}
	Items []Item
)

func (o Order) Validate() error {
	if !validator.IsValidCPF(o.CPF) {
		return errors.New("invalid CPF")
	}
	if err := o.Items.Validate(); err != nil {
		return err
	}
	total := o.Total()
	if total <= 0 {
		return errors.New("the total order price is invalid")
	}
	return nil
}

func (o Order) Total() float64 {
	createdAt := o.CreatedAt
	if createdAt.IsZero() {
		createdAt = time.Now()
	}
	total := o.Items.Total()
	if o.Coupon.Code != "" {
		total = o.Coupon.ApplyDiscount(total, createdAt)
	}
	return total
}

func (os Items) Total() float64 {
	total := float64(0)
	for i := range os {
		total += os[i].total()
	}
	return total
}

func (os Items) Validate() error {
	if len(os) == 0 {
		return errors.New("order without items")
	}
	mapItems := make(map[uint64]bool)
	for _, item := range os {
		if err := item.validate(mapItems); err != nil {
			return err
		}
		mapItems[item.ItemID] = true
	}
	return nil
}

func (o Item) validate(mapItems map[uint64]bool) error {
	if o.Amount < 1 {
		return fmt.Errorf("invalid quantity of %d", o.ItemID)
	}
	if _, ok := mapItems[o.ItemID]; ok {
		return fmt.Errorf("duplicated item %d", o.ItemID)
	}
	return nil
}

func (o Item) total() float64 {
	return o.Price * float64(o.Amount)
}
