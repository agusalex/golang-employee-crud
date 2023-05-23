package models

import (
	"github.com/jinzhu/gorm"
)

type Member struct {
	gorm.Model
	Name             string `gorm:"not null" validate:"required"`
	Type             string `gorm:"type:enum('EMPLOYEE', 'CONTRACTOR');not null" validate:"required,oneof=EMPLOYEE CONTRACTOR"`
	Tags             []Tag  `gorm:"many2many:member_tags;"`
	Role             string `gorm:"" validate:"required_if=Type EMPLOYEE"`
	ContractDuration int    `gorm:"" validate:"required_if=Type CONTRACTOR,gt=0"`
}

/*
func (member *Member) BeforeSave(tx *gorm.DB) (err error) {
	for i, tag := range member.Tags {
		var existingTag Tag
		if err := tx.Where("name = ?", tag.Name).First(&existingTag).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			} else {
				return err
			}
		} else {
			member.Tags[i] = existingTag
		}
	}
	return
}*/
