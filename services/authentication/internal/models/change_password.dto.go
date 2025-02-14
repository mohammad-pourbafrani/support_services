package models

type ChangePasswordDto struct {
	Code        string `json:"code"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}
