package user

import (
	"context"
	"siap_app/internal/app/entity/user"
)

type userUC interface {
	GetListUserAll(ctx context.Context, input user.PaginationUser) ([]user.ResponseUser, int64, error)
	CreateUser(ctx context.Context, input user.RegisterRequest) error
	CreateUserByAdmin(ctx context.Context, input user.RegisterByAdminRequest) error
	LoginUser(ctx context.Context, ipAddress string, input user.LoginRequest, userAgent string) (user.ResponseLogin, error)
	LogoutUser(ctx context.Context, email string) error
	UpdateRoleUser(ctx context.Context, userId int, input user.UpdateRoleRequest) error
	UpdatePasswordUser(ctx context.Context, userId int, input user.UpdatePaswordeRequest) error
	GetUserByEmail(ctx context.Context, email string) (user.ResponseUser, error)
	GetUserById(ctx context.Context, id string) (user.ResponseUser, error)
	UpdateUser(ctx context.Context, userId int, input user.RequestUpdateUser) error
	GetLogLogin(ctx context.Context, input user.PaginationLog) ([]user.ResponseLog, int64, error)
}
