package repository

import (
	"errors"
	"fmt"
	"support_services_authentication/models"
	"time"
)

func (c *authenticationRepository) AddUser(data *models.UserDto) (*models.User, error) {
	query := `INSERT INTO "users" (phone_number, user_role, user_status, created_at) VALUES ($1,$2,$3,NOW()) RETURNING user_id`
	row := c.db.QueryRow(query, data.Phone_number, data.UserRole, data.UserStatus)
	if row.Err() != nil {
		fmt.Println(row.Err())
		return nil, errors.New("internal issue, error code #1008")
	}
	var result int32
	if err := row.Scan(&result); err != nil {
		return nil, errors.New("internal issue, error code #1009")
	}
	return &models.User{
		UserId:       result,
		Phone_number: data.Phone_number,
		UserRole:     data.UserRole,
		UserStatus:   data.UserStatus,
		CreatedAt:    time.Now(),
	}, nil
}
