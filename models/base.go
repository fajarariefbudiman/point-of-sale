package models

import (
	"pos-echo/db"
)

const (
	TaxPercent = 10
)

func GetTaxPercent() float64 {
	return float64(TaxPercent) / 10.0
}

func GetTaxAmount(price float64) float64 {
	return GetTaxPercent() * price
}

func (c *Cart) CalcuCart(cartId string) (*Cart, error) {
	var updatecart, cart Cart
	cond := db.NewDB()
	cartbasetotalprice := 0.0
	carttaxamount := 0.0
	cartdiscountamount := 0.0
	cartgrandtotal := 0.0

	for _, item := range cart.Cart_Items {
		itembasetotal := item.Base_Total
		itemtaxamount := item.Tax_Amount
		itemsubtotaltaxamount := itemtaxamount * float64(item.Quantity)
		itemdiscountamount := item.Discount_Amount
		itemsubtotaldiscountamount := itemdiscountamount * float64(item.Quantity)
		itemsubtotal := item.Sub_Total

		cartbasetotalprice += itembasetotal
		carttaxamount += itemsubtotaltaxamount
		cartdiscountamount += itemsubtotaldiscountamount
		cartgrandtotal += itemsubtotal
	}

	updatecart.Base_Total_Price = cartbasetotalprice
	updatecart.Tax_Amount = carttaxamount
	updatecart.Discount_Amount = cartdiscountamount
	updatecart.Grand_Total = cartgrandtotal

	err := cond.Debug().Where("id = ?", cartId).Model(&Cart{}).Updates(&updatecart).Error

	if err != nil {
		return nil, err
	}

	return &cart, nil

}
