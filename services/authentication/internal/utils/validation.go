package utils

import (
	"support_services_authentication/internal/types"
	"time"
)

func ValidateTokenExpireTime(tokenExpireAt time.Time) *types.Error {
	if time.Now().After(tokenExpireAt) {
		return types.NewBadRequestError("token expired, error code #9029")
	}

	return nil
}
