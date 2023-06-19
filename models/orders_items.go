package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type OrderItem struct {
	Id               string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Orders           Orders
	Order_Id         string
	Product          Products
	Product_Id       string
	Quantity         int
	Base_Price       decimal.Decimal
	Base_Total       decimal.Decimal
	Tax_Amount       decimal.Decimal
	Tax_Percent      decimal.Decimal
	Discount_Amount  decimal.Decimal
	Discount_Percent decimal.Decimal
	Sub_Total        decimal.Decimal
	Sku              string
	Name             string
	Weight           decimal.Decimal
	Created_At       time.Time
	Updated_At       time.Time
}
