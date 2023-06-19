package models

import (
	"pos-echo/db"
	"time"

	"gorm.io/gorm"
)

type Products struct {
	Id                int `gorm:"primary_key"`
	Parent_Id         string
	Users             Users  `gorm:"foreignKey:User_Id"`
	User_Id           string `gorm:"column:User_Id"`
	SKU               string
	Type              string
	Addresses         []Address       `gorm:"foreignKey:User_Id"`
	ProductImages     []ProductImages `gorm:"foreignKey:Product_Id"`
	Categories        []Categories    `gorm:"many2many:product_categories;foreignKey:Id;joinForeignKey:Category_Id;joinReferences:Product_Id"`
	Name              string
	Slug              string
	Price             float64
	Stock             int
	Weight            float64
	Short_Description string
	Description       string
	Status            int
	Created_At        time.Time
	Updated_At        time.Time
	Deleted_At        gorm.DeletedAt
}

func GetProducts() []Products {
	var products []Products
	var err error
	cond := db.NewDB()
	err = cond.Debug().Limit(20).Find(&products).Error
	if err != nil {
		panic("Error Debug")
	}
	return products
}
