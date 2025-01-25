package repository

import (
	"database/sql"
	"support_services_authentication/internal/models"
	"support_services_authentication/internal/types"

	"github.com/redis/go-redis/v9"
)

type (
	AuthenticationRepository interface {
		AddUser(data *models.UserDto) (*models.User, *types.Error)
		AddPassword(data *models.PasswordDto) *types.Error
		SetVerifyCode(data *models.VerifyCodeDto) *types.Error
	}
	authenticationRepository struct {
		db   *sql.DB
		rdDb *redis.Client
	}
)

func NewAuthenticationRepository(db *sql.DB, rdDb *redis.Client) *authenticationRepository {
	return &authenticationRepository{db: db, rdDb: rdDb}
}
