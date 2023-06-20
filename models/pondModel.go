package models

import "gorm.io/gorm"

// User has many CreditCards, UserID is the foreign key

type Pond struct {
	gorm.Model
	PondName string
	FarmID   uint
}
