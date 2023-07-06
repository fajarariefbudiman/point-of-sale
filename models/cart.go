package models

import "pos-echo/db"

type Cart struct {
	Id               string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Cart_Items       []CartItem
	Base_Total_Price float64
	Tax_Amount       float64
	Tax_Percent      float64
	Discount_Amount  float64
	Discount_Percent float64
	Grand_Total      float64
}

func CreateCart(cartId string) (*Cart, error) {
	cond := db.NewDB()
	cart := &Cart{
		Id:               cartId,
		Cart_Items:       []CartItem{},
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

func GetCart(cartId string) (*Cart, error) {
	var cart Cart
	cond := db.NewDB()
	err := cond.Debug().Preload("Cart_Items").Model(&Cart{}).Where("id = ?", cartId).First(&cart).Error
	if err != nil {
		return nil, err
	}

	return &cart, nil
}

func (cart *Cart) AddItem(item CartItem) (*CartItem, error) {
	cond := db.NewDB()
	var exisItem, updateItem CartItem
	var product Products
	err := cond.Debug().Model(Products{}).Where("id = ?", item.Product_Id).First(product).Error
	if err != nil {
		return nil, err
	}
	baseprice := product.Price
	taxamount := GetTaxAmount(baseprice)
	discountamount := 0.0

	err = cond.Debug().Model(CartItem{}).Where("cart_id=?", cart.Id).Where("product_id=?", product.Id).First(&exisItem).Error
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
