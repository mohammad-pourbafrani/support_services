package repository

import (
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