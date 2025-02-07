package models

import "time"

type Token struct {
	AccessToken       string    `json:"access"`
	RefreshToken      string    `json:"refresh"`
	AccessExpireTime  time.Time `json:"access_expire_time"`
	RefreshExpireTime time.Time `json:"refresh_expire_time"`
}
