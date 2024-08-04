package repo_interfaces

import "github.com/shivaraj-shanthaiah/godoc-user/pkg/models"

type UserRepoInterface interface{
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
	FindAllUsers() ([] *models.User, error)
	FindUserByID(id uint) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
}