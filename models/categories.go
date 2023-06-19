package models

import "time"

type Categories struct {
	Id         string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Parent_Id  string
	Sections   []Sections `gorm:"many2many:section_categories;"`
	Section_Id string
	Products   []Products `gorm:"many2many:product_categories;"`
	Name       string
	Slug       string
	Created_At time.Time
	Updated_At time.Time
}
