package category

import "gorm.io/gorm"

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) Create(name string) error {
	return s.db.Create(&Category{Name: name}).Error
}

func (s *Service) List() ([]Category, error) {
	categories := make([]Category, 0)
	err := s.db.Find(&categories).Error
	return categories, err
}
