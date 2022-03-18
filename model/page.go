package model

import (
	"gorm.io/gorm"
)

type Page struct {
	gorm.Model
	Code  string
	Price uint
}
