package models

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	Name    string   `gorm:"unique;not null"`
	Members []Member `gorm:"many2many:member_tags;"`
}
