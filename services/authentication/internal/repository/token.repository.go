package repository

import (
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


