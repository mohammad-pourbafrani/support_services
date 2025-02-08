package repository

import (
	"database/sql"
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

func (c *authenticationRepository) SetVerifyUser(phoneNumber *string) (*models.User, *types.Error) {
	user, selectErr := c.FindUserWithPhoneNumber(phoneNumber)
	if selectErr != nil {
		return nil, selectErr
	}
	query := `UPDATE "users" SET verify = $1 WHERE user_id = $2`
	_, err := c.db.Exec(query, true, user.UserId)
	if err != nil {
		return nil, types.NewInternalError("internal issue, error code #1006")
	}
	user.Verify = true
	return user, nil

}

func (c *authenticationRepository) FindUserWithPhoneNumber(phoneNumber *string) (*models.User, *types.Error) {
	var user models.User
	query := `SELECT user_id, phone_number, user_role, verify, created_at FROM "users" WHERE phone_number = $1`
	err := c.db.QueryRow(query, phoneNumber).Scan(&user.UserId, &user.PhoneNumber, &user.UserRole, &user.Verify, &user.CreatedAt)
	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			return nil, types.NewInternalError("internal issue , error code #1004")
		}
		return nil, types.NewInternalError("internal issue , error code #1005")
	}
	return &user, nil
}

func (c *authenticationRepository) UserExistsByPhoneNumber(phoneNumber *string) (bool, *types.Error) {
	query := `SELECT 1 FROM "users" WHERE phone_number = $1`
	var result int32
	err := c.db.QueryRow(query, phoneNumber).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, types.NewInternalError("internal issue, error code #1005")
	}
	return true, nil
}
