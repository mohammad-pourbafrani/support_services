package models

import "time"

type TokenDto struct {
	AccessToken       string    `json:"access_token"`
	RefreshToken      string    `json:"refresh_token"`
	UserId            int64     `json:"user_id"`
	UserRole          string    `json:"user_role"`
	AccessExpireTime  time.Time `json:"access_expire_time_at"`
	RefreshExpireTime time.Time `json:"refresh_expire_time_at"`
}
