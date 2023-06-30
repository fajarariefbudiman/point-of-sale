package models

type ProductImages struct {
	Id          string   `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Product     Products `gorm:"foreignKey:Products_Id"`
	Products_Id int
	Path        string
	Extra_Large string
	Large       string
	Medium      string
	Small       string
	Created_At  []uint8
	Updated_At  []uint8
}
