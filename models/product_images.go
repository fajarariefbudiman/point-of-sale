package models

import (
	"time"
)

type ProductImages struct {
	Id          string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Product     Products
	Product_Id  string
	Path        string
	Extra_Large string
	Large       string
	Medium      string
	Small       string
	Created_At  time.Time
	Updated_At  time.Time
}
