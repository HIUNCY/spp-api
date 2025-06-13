package repository

import (
	"errors"

	"github.com/hiuncy/spp/models"
	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(user models.User) error
	Login(email, password string) (*models.User, error)
}

type userRepo struct {
	db *gorm.DB
}

// CreateUser implements UserRepo.
func (u *userRepo) CreateUser(user models.User) error {
	panic("unimplemented")
}

// Login implements UserRepo.
func (u *userRepo) Login(email string, password string) (*models.User, error) {
	var user models.User
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, errors.New("password is incorrect")
	}
	return &user, nil
}

func NewUserRepository(db *gorm.DB) UserRepo {
	return &userRepo{db}
}
