package models

import "time"

type Address struct {
	ID          string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	User        Users
	User_Id     string `gorm:"column:user_id"`
	Name        string
	Address1    string
	Address2    string
	Phone       string
	Email       string
	City_Id     string
	Province_Id string
	PostCode    string
	IsPrimary   bool
	Created_At  time.Time
	Updated_At  time.Time
}
