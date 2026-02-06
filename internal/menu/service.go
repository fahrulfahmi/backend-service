package menu

import "gorm.io/gorm"

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) Create(name string, price int64, categoryID uint) error {
	menu := Menu{
		Name:        name,
		Price:       price,
		CategoryID:  categoryID,
		IsAvailable: true,
	}
	return s.db.Create(&menu).Error
}

func (s *Service) List() ([]Menu, error) {
	menus := make([]Menu, 0)
	err := s.db.Find(&menus).Error
	return menus, err
}
