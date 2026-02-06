package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) Create(name, email, password, role string) error {
	if role != "ADMIN" && role != "CASHIER" {
		return errors.New("invalid role")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}

	user := User{
		Name:     name,
		Email:    email,
		Password: string(hash),
		Role:     role,
	}

	return s.db.Create(&user).Error
}
