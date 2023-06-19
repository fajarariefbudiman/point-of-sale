package models

import "time"

type OrderCustomer struct {
	Id          string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	User        Users
	User_Id     string
	Order       Orders
	Order_Id    string
	First_Name  string
	Last_Name   string
	Address1    string
	Address2    string
	Phone       string
	Email       string
	City_Id     string
	Province_Id string
	PostCode    string
	Created_At  time.Time
	Updated_At  time.Time
}
