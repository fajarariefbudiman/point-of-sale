package models

import (
	"pos-echo/db"

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
	Created_At        []uint8
	Updated_At        []uint8
	Deleted_At        gorm.DeletedAt
}

func GetProducts(perpage int, page int) ([]Products, int, error) {
	var products []Products
	var err error
	var count int64

	cond := db.NewDB()

	err = cond.Debug().Model(&Products{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * perpage

	rows, err := cond.Debug().Model(&Products{}).Order("created_at desc").Limit(perpage).Offset(offset).Find(&products).Rows()
	defer rows.Close()
	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var product Products
		err := cond.ScanRows(rows, &product)
		if err != nil {
			return nil, 0, err
		}
		products = append(products)
	}

	return products, int(count), nil
}

func FindProducts(slug string) (*Products, error) {
	var product Products
	cond := db.NewDB()
	err := cond.Debug().Model(&Products{}).Where("slug = ?", slug).First(product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}
