package models

type UserDto struct {
	UserId       int32  `json:"user_id"`
	Phone_number string `json:"phone_number"`
	UserRole     string `json:"user_role"`
	UserStatus   string `json:"user_status"`
}
