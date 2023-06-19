package models

import "time"

type Sections struct {
	Id         string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name       string
	Slug       string
	Created_At time.Time
	Updated_At time.Time
	Categories []Categories `gorm:"many2many:section_categories;"`
}
