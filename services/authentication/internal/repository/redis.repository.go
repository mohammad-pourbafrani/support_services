package repository

import (
	"context"
	"fmt"
	"support_services_authentication/internal/models"
	"support_services_authentication/internal/types"
	"time"
)

func (r *authenticationRepository) SetVerifyCode(data *models.VerifyCodeDto) *types.Error {
	ctx := context.Background()
	_, err := r.rdDb.Set(ctx, data.PhoneNumber, data.Code, time.Millisecond*120).Result()
	if err != nil {
		fmt.Println(err)
		return types.NewInternalError("internal issue , error code #1001")
	}
	return nil
}
