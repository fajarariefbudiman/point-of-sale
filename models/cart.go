package models

import (
	"pos-echo/db"

	"github.com/google/uuid"
)

type Cart struct {
	Id               string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Cart_Items       []Cart_Item
	Base_Total_Price float64
	Tax_Amount       float64
	Tax_Percent      float64
	Discount_Amount  float64
	Discount_Percent float64
	Grand_Total      float64
}

func GetCart(cartId string) (*Cart, error) {
	cond := db.NewDB()
	var cart Cart
	err := cond.Debug().Preload("Cart_Items").Model(&Cart{}).Where("id = ?", cartId).First(&Cart{}).Error
	if err != nil {
		return nil, err
	}

	return &cart, nil
}

func CreateCart() (*Cart, error) {
	cond := db.NewDB()
	cartId := uuid.New().String()
	cart := &Cart{
		Id:               cartId,
		Cart_Items:       []Cart_Item{},
		Base_Total_Price: 0,
		Tax_Amount:       0,
		Tax_Percent:      0,
		Discount_Amount:  0,
		Discount_Percent: 0,
		Grand_Total:      0,
	}

	err := cond.Debug().Create(&cart).Error
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (cart *Cart) AddItem(item Cart_Item) (*Cart_Item, error) {
	cond := db.NewDB()
	var exisItem, updateItem *Cart_Item
	var product *Products
	err := cond.Debug().Model(&Products{}).Where("id = ?", item.Product_Id).First(&product).Error
	if err != nil {
		return nil, err
	}
	baseprice := product.Price
	taxamount := GetTaxAmount(baseprice)
	discountamount := 0.0

	err = cond.Debug().Model(&Cart_Item{}).Where("cart_id = ?", &cart.Id).Where("product_id = ?", &product.Id).First(&exisItem).Error
	if err != nil {
		subtotal := float64(item.Quantity) * (baseprice + taxamount - discountamount)

		item.Cart_Id = cart.Id
		item.Base_Price = product.Price
		item.Base_Total = (baseprice * float64(item.Quantity))
		item.Tax_Percent = GetTaxPercent()
		item.Tax_Amount = taxamount
		item.Discount_Percent = 0
		item.Discount_Amount = discountamount
		item.Sub_Total = subtotal

		err = cond.Debug().Create(&item).Error
		if err != nil {
			return nil, err
		}
		return &item, nil
	}

	updateItem.Quantity = exisItem.Quantity + item.Quantity
	updateItem.Base_Total = baseprice * float64(updateItem.Quantity)
	subTotal := float64(updateItem.Quantity) * (baseprice + taxamount - discountamount)
	updateItem.Sub_Total = subTotal

	err = cond.Debug().First(&exisItem, "id = ?", exisItem.Id).Updates(float64(updateItem.Quantity)).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func UpdateQuantity(itemId string, quantity int) (*Cart_Item, error) {
	cond := db.NewDB()
	var exisItem, updateItem Cart_Item
	err := cond.Debug().Model(Cart_Item{}).Where("id=?", itemId).First(&exisItem).Error
	if err != nil {
		return nil, err
	}
	var product Products
	err = cond.Debug().Model(Products{}).Where("id-?", exisItem.Product_Id).First(&product).Error
	if err != nil {
		return nil, err
	}
	baseprice := product.Price
	taxamount := GetTaxAmount(baseprice)
	discountamount := 0.0
	updateItem.Quantity = quantity
	updateItem.Base_Total = baseprice * float64(updateItem.Quantity)
	subTotal := float64(updateItem.Quantity) * (baseprice + taxamount - discountamount)
	updateItem.Sub_Total = subTotal

	err = cond.Debug().First(&exisItem, "id=?", exisItem.Id).Updates(&updateItem).Error
	if err != nil {
		return nil, err
	}
	return &exisItem, err
}
