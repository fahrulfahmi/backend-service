package auth

import (
	"errors"

	"backend-service/internal/user"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) Login(email, password string) (*user.User, error) {
	var u user.User

	if err := s.db.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(u.Password),
		[]byte(password),
	); err != nil {
		return nil, errors.New("invalid password")
	}

	return &u, nil
}
