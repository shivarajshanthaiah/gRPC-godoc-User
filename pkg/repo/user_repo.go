package repo

import (
	"github.com/shivaraj-shanthaiah/godoc-user/pkg/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (u *UserRepository) CreateUser(user *models.User) error {
	if err := u.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) FindUserByID(id uint) (*models.User, error) {
	var user *models.User
	err := u.DB.First(&user, id).Error
	return user, err
}

func (u *UserRepository) UpdateUser(user *models.User) error {
	err := u.DB.Save(&user).Error
	return err
}

func (u *UserRepository) DeleteUser(id uint) error {
	err := u.DB.Delete(&models.User{}, id).Error
	return err
}

func (u *UserRepository) FindUserByEmail(email string) (*models.User, error) {
	var user *models.User
	err := u.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

func (u *UserRepository) FindAllUsers() ([]*models.User, error) {
	var users []*models.User
	err := u.DB.Find(&users).Error
	return users, err
}
