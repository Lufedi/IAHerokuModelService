package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Rol string
	Email string
}