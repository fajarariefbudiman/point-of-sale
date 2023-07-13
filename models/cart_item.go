package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart_Item struct {
	Id               string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Cart             Cart
	Cart_Id          string
	Product          Products
	Product_Id       string
	Quantity         int
	Base_Price       float64
	Base_Total       float64
	Tax_Amount       float64
	Tax_Percent      float64
	Discount_Amount  float64
	Discount_Percent float64
	Sub_Total        float64
	Created_At       []uint8
	Updated_At       []uint8
}

func (c *Cart_Item) BeforeCreate(tx *gorm.DB) error {
	if c.Id == "" {
		c.Id = uuid.New().String()
	}

	return nil
}
