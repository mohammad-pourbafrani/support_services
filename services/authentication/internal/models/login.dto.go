package models

type LogInDto struct {
	PhoneNumber  string       `json:"phone_number"`
	Password     string       `json:"password"`
	SignInMethod SignInMethod `json:"sign_in_method"`
}

type SignInMethod int

const (
	PASSWORD SignInMethod = iota
	VERIFY_CODE
)
