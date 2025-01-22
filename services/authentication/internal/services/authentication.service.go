package services

import (
	"support_services_authentication/internal/repository"
	"support_services_authentication/internal/models"
)

type (
	AuthenticationService interface {
		AddUser(data *models.UserDto) *models.User
	}

	authenticationService struct {
		repository repository.AuthenticationRepository
	}
)

func NewAuthenticationService(repository repository.AuthenticationRepository) AuthenticationService {
	return &authenticationService{repository: repository}
}

func (c *authenticationService) AddUser(data *models.UserDto) *models.User {
	res, err := c.repository.AddUser(data)
	if err != nil {
		return nil
	} else {
		return res
	}
}
