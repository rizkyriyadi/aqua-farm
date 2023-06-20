package models

import "gorm.io/gorm"

type Farm struct {
	gorm.Model
	FarmName string
	Pond     []Pond `gorm:"foreignKey:farm_id"`
}
