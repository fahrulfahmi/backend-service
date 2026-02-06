package order

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Code        string
	Status      string
	TotalAmount int64
	Items       []OrderItem
}

type OrderItem struct {
	gorm.Model
	OrderID  uint
	MenuID   uint
	MenuName string
	Price    int64
	Qty      int
	Subtotal int64
}
