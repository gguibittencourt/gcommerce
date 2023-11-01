package repositories

import (
	"database/sql"
	"time"

	"github.com/gguibittencourt/gcommerce/app/coupon"
	"github.com/gguibittencourt/gcommerce/app/freight"
	"github.com/gguibittencourt/gcommerce/app/order"
	"github.com/gguibittencourt/gcommerce/internal/database"
)

type freightModel struct {
	database.Model
	OrderID  uint64
	Code     string
	Price    float64
	Duration time.Duration
	ETA      time.Time
}

type itemModel struct {
	database.Model
	OrderID   uint64
	ProductID uint64
	Amount    uint16
	Price     float64
}

type orderModel struct {
	database.Model
	CPF      string
	Code     string
	Status   string
	Total    float64
	CouponID sql.NullInt64
	Freight  freightModel `gorm:"foreignKey:OrderID;references:ID"`
	Items    []itemModel  `gorm:"foreignKey:OrderID;references:ID"`
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

func toOrderModel(o order.Order) orderModel {
	return orderModel{
		CPF:      o.CPF,
		Code:     o.Code,
		Status:   o.Status,
		Total:    o.Total(),
		CouponID: sql.NullInt64{Int64: int64(o.Coupon.CouponID), Valid: o.Coupon.CouponID != 0},
		Freight:  toFreightModel(o.Freight),
		Items:    toItemsModel(o.Items),
	}
}

func toItemsModel(items order.Items) []itemModel {
	models := make([]itemModel, len(items))
	for i := range items {
		models[i] = toItemModel(items[i])
	}
	return models
}

func toItemModel(item order.Item) itemModel {
	return itemModel{
		ProductID: item.ProductID,
		Amount:    uint16(item.Amount),
		Price:     item.Price,
	}
}

func toFreightModel(f freight.Freight) freightModel {
	return freightModel{
		Code:     f.Code,
		Price:    f.Price,
		Duration: f.DurationInMin,
		ETA:      f.ETA,
	}
}

func (o orderModel) toOrder() order.Order {
	items := make(order.Items, len(o.Items))
	for i := range o.Items {
		items[i] = o.Items[i].toItem()
	}
	return order.Order{
		OrderID: o.ID,
		CPF:     o.CPF,
		Coupon: coupon.Coupon{
			CouponID: uint64(o.CouponID.Int64),
		},
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

func (f freightModel) toFreight() freight.Freight {
	return freight.Freight{
		FreightID:     f.ID,
		Price:         f.Price,
		DurationInMin: f.Duration,
		ETA:           f.ETA,
	}
}
