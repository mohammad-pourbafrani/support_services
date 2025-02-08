package services

import (
	"fmt"
	"support_services_authentication/internal/models"
	"support_services_authentication/internal/repository"
	"support_services_authentication/internal/types"
	"support_services_authentication/internal/utils"
)

type (
	AuthenticationService interface {
		RegisterUser(data *models.UserDto) (*models.User, *types.Error)
		VerifyRegisterUser(data *models.VerifyCodeDto) (*models.User, *types.Error)
	}

	authenticationService struct {
		repository repository.AuthenticationRepository
	}
)

func NewAuthenticationService(repository repository.AuthenticationRepository) AuthenticationService {
	return &authenticationService{repository: repository}
}

func (c *authenticationService) RegisterUser(data *models.UserDto) (*models.User, *types.Error) {

	if userExist, existErr := c.repository.UserExistsByPhoneNumber(&data.PhoneNumber); existErr != nil {
		return nil, existErr
	} else if userExist {
		return nil, types.NewBadRequestError("user exist , please sign in or reset password")
	}

	res, userErr := c.repository.AddUser(data)
	if userErr != nil {
		return nil, userErr
	} else if passErr := c.repository.AddPassword(&models.PasswordDto{UserId: res.UserId, Password: data.Password}); passErr != nil {
		return nil, passErr
	}
	var (
		check      bool = true
		verifyCode int32
	)

	for check {
		code, randErr := utils.NextRandomInt32(120001, 510001)
		if randErr != nil {
			return nil, types.NewInternalError(randErr.Error())
		}

		existCode, err := c.repository.CheckExistCode(fmt.Sprintf("%d", code))
		if err != nil {
			return nil, err
		}
		if !existCode {
			verifyCode = code
			check = false
		}
	}
	//TODO: send verify code to user
	fmt.Printf("verify code : %v", verifyCode)
	redisError := c.repository.SetVerifyCode(&models.VerifyCodeDto{PhoneNumber: res.PhoneNumber, Code: fmt.Sprintf("%d", verifyCode)})
	if redisError != nil {
		return nil, redisError
	}
	return res, nil
}

func (c *authenticationService) VerifyRegisterUser(data *models.VerifyCodeDto) (*models.User, *types.Error) {

	user, err := c.repository.FindUserWithPhoneNumber(&data.PhoneNumber)
	if err != nil {
		return nil, err
	}
	if user.Verify {
		return nil, types.NewBadRequestError("user exist! please sign in")
	}

	res, redisError := c.repository.GetVerifyCode(data)
	if redisError != nil {
		return nil, redisError
	}

	if *res == data.PhoneNumber {
		user, err := c.repository.SetVerifyUser(&data.PhoneNumber)
		if err != nil {
			return nil, err
		}
		return user, nil
	} else {
		return nil, types.NewBadRequestError("your verify code not exist")
	}

}
