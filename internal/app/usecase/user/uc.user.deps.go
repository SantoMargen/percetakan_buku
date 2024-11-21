package user

import (
	"context"
	loglogin "siap_app/internal/app/entity/log_login"
	"siap_app/internal/app/entity/user"
)

type userRepo interface {
	GetListUserAll(ctx context.Context, input user.PaginationUser) ([]user.ResponseUser, int64, error)
	CreateUser(ctx context.Context, input user.RegisterRequest) error
	CreateUserByAdmin(ctx context.Context, input user.RegisterByAdminRequest) error
	GetUserByEmail(ctx context.Context, email string) (user.ResponseUser, error)
	GetUserById(ctx context.Context, email string) (user.ResponseUser, error)
	UpdateRoleUser(ctx context.Context, id, userId int, role string) error
	UpdatePasswordUser(ctx context.Context, userId int, password string) error
	UpdateUser(ctx context.Context, userId int, input user.RequestUpdateUser) error
	GetLogLogin(ctx context.Context, input user.PaginationLog) ([]user.ResponseLog, int64, error)
}

type redisRepo interface {
	SaveTokenInRedis(key, data string) error
	DeleteTokenRedis(key string) error
	GetRedisData(key string) (string, error)
	SaveAccountFreeze(key, data string) error
}

type logLoginRepo interface {
	CreateLogLogin(ctx context.Context, logLogin loglogin.LogloginRequest) error
	CheckTableAlreadyExist(ctx context.Context) bool
	CreateTableLogLogin(ctx context.Context) error
	GetLastLogLoginByEmail(ctx context.Context, email string) (*loglogin.LogloginResponse, error)
	UpdateLogLogin(ctx context.Context, logLogin loglogin.LogloginRequest) error
}
