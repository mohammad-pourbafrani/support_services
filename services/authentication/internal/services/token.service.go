package services

import (
	"support_services_authentication/internal/models"
	"support_services_authentication/internal/repository"
	"support_services_authentication/internal/types"
	"support_services_authentication/internal/utils"
	"time"
)

type (
	TokenService interface {
		AddToken(user *models.User) (*models.Token, *types.Error)
		RenewToken(data *models.RenewTokenDto) (*models.Token, *types.Error)
	}
	tokenService struct {
		repository repository.AuthenticationRepository
	}
)

func NewTokenService(repository repository.AuthenticationRepository) TokenService {
	return &tokenService{repository: repository}
}

func (s *tokenService) AddToken(user *models.User) (*models.Token, *types.Error) {

	accessToken, accessErr := utils.NextAlphanumeric(40)
	if accessErr != nil {
		return nil, types.NewInternalError("internal issue , error code #1008")
	}
	refreshToken, refreshErr := utils.NextAlphanumeric(40)
	if refreshErr != nil {
		return nil, types.NewInternalError("internal issue , error code #1008")
	}
	token := &models.TokenDto{
		AccessToken:       accessToken,
		RefreshToken:      refreshToken,
		UserId:            user.UserId,
		UserRole:          user.UserRole,
		AccessExpireTime:  time.Now().Add(time.Minute * 15),
		RefreshExpireTime: time.Now().AddDate(0, 0, 5),
	}
	tokenErr := s.repository.AddToken(token)
	if tokenErr != nil {
		return nil, tokenErr
	}

	return &models.Token{
		AccessToken:       accessToken,
		RefreshToken:      refreshToken,
		AccessExpireTime:  token.AccessExpireTime,
		RefreshExpireTime: token.RefreshExpireTime,
	}, nil
}

func (c *tokenService) RenewToken(data *models.RenewTokenDto) (*models.Token, *types.Error) {

	token, tokenErr := c.repository.GetTokenByAccessTokenAndRefreshToken(&data.AccessToken, &data.RefreshToken)
	if tokenErr != nil {
		return nil, tokenErr
	}

	refreshErr := utils.ValidateTokenExpireTime(token.RefreshExpireTime)

	if refreshErr != nil {
		return nil, refreshErr
	}

	deleteTokenErr := c.repository.DeleteTokenWithAccessToken(token.AccessToken)
	if deleteTokenErr != nil {
		return nil, deleteTokenErr
	}

	user, _, err := c.repository.FindUserWithUserId(&token.UserId)
	if err != nil {
		return nil, err
	}

	newToken, addTokenErr := c.AddToken(user)
	if addTokenErr != nil {
		return nil, addTokenErr
	}

	return newToken, nil

}
