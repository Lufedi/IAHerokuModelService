package model

import (
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	User User
	UserID int
	Description string
	Name string
	Location string
}
