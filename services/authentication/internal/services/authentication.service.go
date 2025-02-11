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
		RegisterUser(data *models.UserDto) *types.Error
		VerifyRegisterUser(data *models.VerifyCodeDto) (*models.User, *types.Error)
		LogIn(data *models.LogInDto) (*models.User, *types.Error)
		LogInWithVerifyCode(data *models.VerifyCodeDto) (*models.User, *types.Error)
	}

	authenticationService struct {
		repository repository.AuthenticationRepository
	}
)

func NewAuthenticationService(repository repository.AuthenticationRepository) AuthenticationService {
	return &authenticationService{repository: repository}
}

func (c *authenticationService) RegisterUser(data *models.UserDto) *types.Error {

	if user, exist, existErr := c.repository.FindUserWithPhoneNumber(&data.PhoneNumber); existErr != nil {
		return existErr
	} else if user != nil && user.Verify {
		return types.NewBadRequestError("user exist , please sign in or reset password")
	} else if !exist {
		res, userErr := c.repository.AddUser(data)
		if userErr != nil {
			return userErr
		} else if passErr := c.repository.AddPassword(&models.PasswordDto{UserId: res.UserId, Password: data.Password}); passErr != nil {
			return passErr
		}
	}

	var (
		check      bool = true
		verifyCode int32
	)

	for check {
		code, randErr := utils.NextRandomInt32(120001, 510001)
		if randErr != nil {
			return types.NewInternalError(randErr.Error())
		}

		existCode, err := c.repository.CheckExistCode(fmt.Sprintf("%d", code))
		if err != nil {
			return err
		}
		if !existCode {
			verifyCode = code
			check = false
		}
	}
	//TODO: send verify code to user
	fmt.Printf("verify code : %v", verifyCode)
	redisError := c.repository.SetVerifyCode(&models.VerifyCodeDto{PhoneNumber: data.PhoneNumber, Code: fmt.Sprintf("%d", verifyCode)})
	if redisError != nil {
		return redisError
	}
	return nil
}

func (c *authenticationService) VerifyRegisterUser(data *models.VerifyCodeDto) (*models.User, *types.Error) {

	user, _, err := c.repository.FindUserWithPhoneNumber(&data.PhoneNumber)
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

func (c *authenticationService) LogIn(data *models.LogInDto) (*models.User, *types.Error) {

	user, exist, err := c.repository.FindUserWithPhoneNumber(&data.PhoneNumber)

	if err != nil {
		return nil, err
	}

	if exist && !user.Verify {
		return nil, types.NewBadRequestError("user not verified , please verify user")
	} else if !exist {
		return nil, types.NewBadRequestError("user not exist , please sign up")
	}

	if data.SignInMethod == models.PASSWORD {
		pass, _, passErr := c.repository.GetPasswordWithUserId(&user.UserId)
		if passErr != nil {
			return nil, passErr
		}
		if data.Password == pass.Password {
			return user, nil
		} else {
			return nil, types.NewBadRequestError("the password is wrong")
		}
	}

	if data.SignInMethod == models.VERIFY_CODE {
		var (
			check      bool = true
			verifyCode int32
		)

		for check {
			code, randErr := utils.NextRandomInt32(510002, 990099)
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
		redisError := c.repository.SetVerifyCode(&models.VerifyCodeDto{PhoneNumber: data.PhoneNumber, Code: fmt.Sprintf("%d", verifyCode)})
		if redisError != nil {
			return nil, redisError
		}
	}

	return user, nil

}

func (c *authenticationService) LogInWithVerifyCode(data *models.VerifyCodeDto) (*models.User, *types.Error) {

	user, exist, err := c.repository.FindUserWithPhoneNumber(&data.PhoneNumber)

	if err != nil {
		return nil, err
	}

	if exist && !user.Verify {
		return nil, types.NewBadRequestError("user not verified , please verify user")
	} else if !exist {
		return nil, types.NewBadRequestError("user not exist , please sign up")
	}

	res, redisError := c.repository.GetVerifyCode(data)
	if redisError != nil {
		return nil, redisError
	}

	if *res == data.PhoneNumber {
		return user, nil
	} else {
		return nil, types.NewBadRequestError("your verify code not exist")
	}

}
