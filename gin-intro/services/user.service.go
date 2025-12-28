package services

import (
	"github.com/nitesh111sinha/apis/models"
)

type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
	ResetUser(*string, *models.User) error
}


