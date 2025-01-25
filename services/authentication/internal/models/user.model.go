package models

import "time"

type User struct {
	UserId      int64     `json:"user_id"`
	PhoneNumber string    `json:"phone_number"`
	UserRole    string    `json:"user_role"`
	Verify      bool      `json:"verify"`
	CreatedAt   time.Time `json:"created_at"`
}
