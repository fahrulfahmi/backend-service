package order

import (
	"fmt"
	"time"

	"backend-service/internal/menu"

	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

type CreateOrderItem struct {
	MenuID uint `json:"menu_id"`
	Qty    int  `json:"qty"`
}

func (s *Service) Create(items []CreateOrderItem) (*Order, error) {
	var total int64
	var orderItems []OrderItem

	for _, item := range items {
		var m menu.Menu
		if err := s.db.First(&m, item.MenuID).Error; err != nil {
			return nil, err
		}

		subtotal := m.Price * int64(item.Qty)
		total += subtotal

		orderItems = append(orderItems, OrderItem{
			MenuID:   m.ID,
			MenuName: m.Name,
			Price:    m.Price,
			Qty:      item.Qty,
			Subtotal: subtotal,
		})
	}

	order := Order{
		Code:        fmt.Sprintf("ORD-%d", time.Now().Unix()),
		Status:      "NEW",
		TotalAmount: total,
		Items:       orderItems,
	}

	err := s.db.Create(&order).Error
	return &order, err
}

func (s *Service) GetByID(id uint) (*Order, error) {
	var order Order
	err := s.db.Preload("Items").First(&order, id).Error
	return &order, err
}
