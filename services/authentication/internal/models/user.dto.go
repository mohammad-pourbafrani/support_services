package models

type UserDto struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	UserRole    string `json:"user_role"`
	Verify      bool   `json:"verify"`
}
