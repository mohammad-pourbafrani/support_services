package services

import (
	"support_services_authentication/internal/models"
	"support_services_authentication/internal/repository"
	"support_services_authentication/internal/types"
	"support_services_authentication/internal/utils"
)

type (
	AuthenticationService interface {
		RegisterUser(data *models.UserDto) (*models.User, *types.Error)
	}

	authenticationService struct {
		repository repository.AuthenticationRepository
	}
)

func NewAuthenticationService(repository repository.AuthenticationRepository) AuthenticationService {
	return &authenticationService{repository: repository}
}

func (c *authenticationService) RegisterUser(data *models.UserDto) (*models.User, *types.Error) {
	res, userErr := c.repository.AddUser(data)
	if userErr != nil {
		return nil, userErr
	} else if passErr := c.repository.AddPassword(&models.PasswordDto{UserId: res.UserId, Password: data.Password}); passErr != nil {
		return nil, passErr
	}
	code, randErr := utils.NextRandomInt32(120001, 510001)
	if randErr != nil {
		return nil, types.NewInternalError(randErr.Error())
	}
	println("vevrify code %v ", code)
	redisEror := c.repository.SetVerifyCode(&models.VerifyCodeDto{PhoneNumber: res.PhoneNumber, Code: code})
	if redisEror != nil {
		return nil, redisEror
	}
	return res, nil
}
