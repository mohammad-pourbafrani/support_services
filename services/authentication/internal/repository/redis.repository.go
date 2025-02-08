package repository

import (
	"context"
	"fmt"
	"support_services_authentication/internal/models"
	"support_services_authentication/internal/types"
	"time"

	"github.com/redis/go-redis/v9"
)

func (r *authenticationRepository) SetVerifyCode(data *models.VerifyCodeDto) *types.Error {
	ctx := context.Background()
	_, err := r.rdDb.Set(ctx, data.Code, data.PhoneNumber, time.Second*120).Result()
	if err != nil {
		fmt.Println(err)
		return types.NewInternalError("internal issue , error code #1001")
	}
	return nil
}

func (r *authenticationRepository) CheckExistCode(data string) (bool, *types.Error) {
	ctx := context.Background()
	err := r.rdDb.Get(ctx, data)
	if err != nil {
		if err.Err() == redis.Nil {
			return false, nil
		}
		fmt.Println(err)
		return false, types.NewInternalError("internal issue , error code #1003")
	}
	return true, nil
}

func (r *authenticationRepository) GetVerifyCode(data *models.VerifyCodeDto) (*string, *types.Error) {
	ctx := context.Background()
	val, err := r.rdDb.GetDel(ctx, data.Code).Result()
	if err != nil {
		fmt.Println(err)
		return nil, types.NewInternalError("internal issue , error code #1003")
	}
	return &val, nil
}
