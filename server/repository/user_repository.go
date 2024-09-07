package repository

import (
	"go-react-todo/server/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *models.User, email string) error
	CreateUser(user *models.User) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetUserByEmail(user *models.User, email string) error {
	if err := r.db.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) CreateUser(user *models.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}
