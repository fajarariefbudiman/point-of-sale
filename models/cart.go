package models

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
