package repository

import (
	"errors"

	"github.com/hiuncy/spp/models"
	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(user *models.User) (int, error)
	Login(email, password string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	DeleteUser(idUser int) error
}

type userRepo struct {
	db *gorm.DB
}

// GetUserByEmail implements UserRepo.
func (u *userRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := u.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser implements UserRepo.
func (u *userRepo) CreateUser(user *models.User) (int, error) {
	if err := u.db.Create(user).Error; err != nil {
		return 0, err
	}
	return user.IdUser, nil
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

func (u *userRepo) DeleteUser(idUser int) error {
	if err := u.db.Delete(&models.User{}, idUser).Error; err != nil {
		return err
	}
	return nil
}

func NewUserRepository(db *gorm.DB) UserRepo {
	return &userRepo{db}
}
