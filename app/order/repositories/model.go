package repositories

import (
	"time"

	"github.com/gguibittencourt/gcommerce/app/freight"
	"github.com/gguibittencourt/gcommerce/app/order"
	"github.com/gguibittencourt/gcommerce/internal/database"
)

type couponModel struct {
	database.Model
	Code           string
	ExpirationDate time.Time
	Percentage     float64
}

type freightModel struct {
	database.Model
	Code          string
	Price         float64
	DurationInMin time.Duration
	ETA           time.Time
}

type itemModel struct {
	database.Model
	ProductID uint64
	Amount    uint16
	Price     float64
}

type orderModel struct {
	database.Model
	CPF     string
	Code    string
	Status  string
	Total   float64
	Coupon  couponModel  `gorm:"foreignKey:CouponID;references:ID"`
	Freight freightModel `gorm:"foreignKey:FreightID;references:ID"`
	Items   []itemModel  `gorm:"foreignKey:ItemID;references:ID"`
}

func (orderModel) TableName() string {
	return "order"
}

func (itemModel) TableName() string {
	return "item"
}

func (freightModel) TableName() string {
	return "freight"
}

func (couponModel) TableName() string {
	return "coupon"
}

func (o orderModel) toOrder() order.Order {
	items := make(order.Items, len(o.Items))
	for i := range o.Items {
		items[i] = o.Items[i].toItem()
	}
	return order.Order{
		OrderID:   o.ID,
		CPF:       o.CPF,
		Coupon:    o.Coupon.toCoupon(),
		Items:     items,
		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,
		Freight:   o.Freight.toFreight(),
	}
}

func (i itemModel) toItem() order.Item {
	return order.Item{
		ItemID:    i.ID,
		ProductID: i.ProductID,
		Amount:    uint32(i.Amount),
		Price:     i.Price,
	}
}
func (c couponModel) toCoupon() order.Coupon {
	return order.Coupon{
		CouponID:       c.ID,
		Code:           c.Code,
		Percentage:     c.Percentage,
		ExpirationDate: c.ExpirationDate,
	}
}

func (f freightModel) toFreight() freight.Freight {
	return freight.Freight{
		FreightID:     f.ID,
		Price:         f.Price,
		DurationInMin: f.DurationInMin,
		ETA:           f.ETA,
	}
}
