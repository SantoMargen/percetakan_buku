package user

import (
	"context"
	"siap_app/internal/app/entity/user"
)

type userUC interface {
	CreateUser(ctx context.Context, input user.RegisterRequest) error
	LoginUser(ctx context.Context, input user.LoginRequest) (user.ResponseLogin, error)
}
