package models

import (
	"time"

	"gorm.io/gorm"

	"github.com/shopspring/decimal"
)

type Orders struct {
	Id                    string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	User_Id               string
	Users                 Users
	OrderItems            []OrderItem
	OrderCustomer         *OrderCustomer
	Code                  string
	Status                int
	Order_Date            time.Time
	Payment_Due           time.Time
	Payment_Status        string
	Payment_Token         string
	Base_Total_Price      decimal.Decimal
	Tax_Amount            decimal.Decimal
	Tax_Percent           decimal.Decimal
	Discount_Amount       decimal.Decimal
	Discount_Percent      decimal.Decimal
	Shipping_Cost         decimal.Decimal
	Grand_Total           decimal.Decimal
	Note                  string
	Shipping_Courier      string
	Shipping_Service_Name string
	Approved_By           string
	Approved_At           time.Time
	Cancelled_By          string
	Cancelled_At          time.Time
	Cancellation_Note     string
	Created_At            time.Time
	Updated_At            time.Time
	Deleted_At            gorm.DeletedAt
}
