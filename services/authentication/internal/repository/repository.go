package repository

import (
	"database/sql"
	"support_services_authentication/models"
)

type (
	AuthenticationRepository interface {
		AddUser(data *models.UserDto) (*models.User, error)
	}
	authenticationRepository struct {
		db *sql.DB
	}
)

func NewAuthenticationRepository(db *sql.DB) *authenticationRepository {
	return &authenticationRepository{db: db}
}
