package models

import "time"

type User struct {
	UserId       int32     `json:"user_id"`
	Phone_number string    `json:"phone_number"`
	UserRole     string    `json:"user_role"`
	UserStatus   string    `json:"user_status"`
	CreatedAt    time.Time `json:"created_at"`
}
