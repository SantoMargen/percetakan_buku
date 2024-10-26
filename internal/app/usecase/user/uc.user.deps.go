package user

import (
	"context"
	"siap_app/internal/app/entity/user"
)

type userRepo interface {
	CreateUser(ctx context.Context, input user.RegisterRequest) error
	GetUserByEmail(ctx context.Context, email string) (user.ResponseUser, error)
}
