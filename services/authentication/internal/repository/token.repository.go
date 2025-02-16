package repository

import (
	"database/sql"
	"fmt"
	"support_services_authentication/internal/models"
	"support_services_authentication/internal/types"
)

func (c *authenticationRepository) AddToken(token *models.TokenDto) *types.Error {
	query := `INSERT INTO "tokens" (access_token, refresh_token, user_id, user_role, created_at, access_token_expire_at, refresh_token_expire_at ) VALUES ($1,$2,$3,$4,NOW(),$5,$6)`
	row := c.db.QueryRow(query, token.AccessToken, token.RefreshToken, token.UserId, token.UserRole, token.AccessExpireTime, token.RefreshExpireTime)
	if row.Err() != nil {
		fmt.Println(row.Err())
		return types.NewInternalError("internal issue , error code #1007")
	}
	return nil
}

func (c *authenticationRepository) DeleteTokensWithUserId(userId *int64) *types.Error {
	query := `DELETE FROM "tokens" WHERE user_id=$1`
	row := c.db.QueryRow(query, userId)
	if row.Err() != nil {
		fmt.Println(row.Err())
		return types.NewInternalError("internal issue , error code #1011")
	}
	return nil
}

func (c *authenticationRepository) GetTokenByAccessTokenAndRefreshToken(accessToken *string, refreshToken *string) (*models.TokenDto, *types.Error) {
	query := `SELECT access_token, refresh_token, user_id, user_role , access_token_expire_at, refresh_token_expire_at FROM "tokens" WHERE access_token = $1 AND refresh_token = $2`
	var token models.TokenDto
	err := c.db.QueryRow(query, accessToken, refreshToken).Scan(&token.AccessToken, &token.RefreshToken, &token.UserId, &token.UserRole, &token.AccessExpireTime, &token.RefreshExpireTime)
	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("token not found, error code #1012")
		}
		return nil, types.NewInternalError("internal issue, error code #1035")
	}
	return &token, nil
}

func (c *authenticationRepository) DeleteTokenWithAccessToken(accessToken string) *types.Error {
	query := `DELETE FROM "tokens" WHERE access_token=$1`
	row := c.db.QueryRow(query, accessToken)
	if row.Err() != nil {
		fmt.Println(row.Err())
		return types.NewInternalError("internal issue , error code #1014")
	}
	return nil
}
