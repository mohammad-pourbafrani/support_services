package repository

import (
	"fmt"
	"support_services_authentication/internal/models"
	"support_services_authentication/internal/types"
	"time"
)

func (c *authenticationRepository) AddUser(data *models.UserDto) (*models.User, *types.Error) {
	query := `INSERT INTO "users" (phone_number, user_role,verify, created_at) VALUES ($1,$2,$3,NOW()) RETURNING user_id`
	row := c.db.QueryRow(query, data.PhoneNumber, data.UserRole, data.Verify)
	if row.Err() != nil {
		fmt.Println(row.Err())
		return nil, types.NewInternalError("internal issue , error code #1001")
	}
	var result int64
	if err := row.Scan(&result); err != nil {
		return nil, types.NewInternalError("internal issue , error code #1002")
	}
	return &models.User{
		UserId:      result,
		PhoneNumber: data.PhoneNumber,
		UserRole:    data.UserRole,
		Verify:      data.Verify,
		CreatedAt:   time.Now(),
	}, nil
}
