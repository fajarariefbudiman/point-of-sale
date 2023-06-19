package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Payment struct {
	Id           string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Order        Orders
	Order_Id     string
	Number       string
	Amount       decimal.Decimal
	Method       string
	Status       string
	Token        string
	Payloads     string
	Payment_Type string
	Va_Number    string
	Biller_Code  string
	Bill_Key     string
	Created_At   time.Time
	Updated_At   time.Time
	Deleted_At   gorm.DeletedAt
}
