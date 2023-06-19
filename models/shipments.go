package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Shipment struct {
	Id             string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	User           Users
	User_Id        string
	Order          Orders
	Order_Id       string
	Track_Number   string
	Status         string
	Total_Quantity int
	Total_Weight   decimal.Decimal
	First_Name     string
	Last_Name      string
	City_Id        string
	Province_Id    string
	Address1       string
	Address2       string
	Phone          string
	Email          string
	PostCode       string
	Shipped_By     string
	Shipped_At     time.Time
	Created_At     time.Time
	Updated_At     time.Time
	Deleted_At     gorm.DeletedAt
}
