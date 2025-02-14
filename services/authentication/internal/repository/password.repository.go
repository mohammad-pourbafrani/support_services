package repository

import (
	"database/sql"
	"fmt"
	"support_services_authentication/internal/models"
	"support_services_authentication/internal/types"
)

func (c *authenticationRepository) AddPassword(data *models.PasswordDto) *types.Error {
	query := `INSERT INTO "passwords" (user_id, password) VALUES ($1,$2)`
	row := c.db.QueryRow(query, data.UserId, data.Password)
	if row.Err() != nil {
		fmt.Println(row.Err())
		return types.NewInternalError("internal issue , error code #1001")
	}
	return nil
}

func (c *authenticationRepository) GetPasswordWithUserId(userId *int64) (*models.Password, bool, *types.Error) {
	var pass models.Password
	query := `SELECT password FROM "passwords" WHERE user_id = $1`
	err := c.db.QueryRow(query, userId).Scan(&pass.Password)
	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			return nil, false, nil
		}
		return nil, false, types.NewInternalError("internal issue , error code #1004")
	}
	return &pass, true, nil
}

func (c *authenticationRepository) UpdatePassword(data *models.PasswordDto) *types.Error {
	query := `UPDATE "passwords" SET password = $1 WHERE user_id = $2`
	_, err := c.db.Exec(query, data.Password, data.UserId)
	if err != nil {
		return types.NewInternalError("internal issue, error code #1010")
	}
	return nil
}
