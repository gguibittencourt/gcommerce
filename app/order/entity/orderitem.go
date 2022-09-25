package entity

import (
	"errors"
	"fmt"
)

type (
	OrderItem struct {
		ItemID uint64
		Amount uint32
		Price  float64
	}
	OrderItems []OrderItem
)

func (os OrderItems) Total() float64 {
	total := float64(0)
	for i := range os {
		total += os[i].total()
	}
	return total
}

func (os OrderItems) Validate() error {
	if len(os) == 0 {
		return errors.New("order without items")
	}
	mapItems := make(map[uint64]bool)
	for _, item := range os {
		err := item.validate(mapItems)
		if err != nil {
			return err
		}
		mapItems[item.ItemID] = true
	}
	return nil
}

func (o OrderItem) validate(mapItems map[uint64]bool) error {
	if o.Amount < 1 {
		return fmt.Errorf("invalid quantity of %d", o.ItemID)
	}
	if _, ok := mapItems[o.ItemID]; ok {
		return fmt.Errorf("duplicated item %d", o.ItemID)
	}
	return nil
}

func (o OrderItem) total() float64 {
	return o.Price * float64(o.Amount)
}
