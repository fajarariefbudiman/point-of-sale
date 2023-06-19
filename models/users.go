package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id             string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	User_Id        string
	Addresses      []Address `gorm:"foreignKey:User_Id"`
	First_Name     string
	Last_Name      string
	Email          string
	Password       string
	Remember_Token string
	Created_At     time.Time
	Updated_At     time.Time
	Deleted_At     gorm.DeletedAt
}
