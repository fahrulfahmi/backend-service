package menu

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Name        string
	Price       int64
	CategoryID  uint
	IsAvailable bool
}
