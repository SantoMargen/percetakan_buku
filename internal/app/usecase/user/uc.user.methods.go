package user

import (
	"context"
	"encoding/json"
	"fmt"
	"siap_app/internal/app/entity"
	loglogin "siap_app/internal/app/entity/log_login"
	"siap_app/internal/app/entity/user"
	"siap_app/internal/app/helpers"
	"strings"
	"time"
)

func (uc *UseCase) CreateUser(ctx context.Context, input user.RegisterRequest) error {
	role := "AUTHOR"
	hasPass, err := helpers.HashPassword(input.Password)
	if err != nil {
		return fmt.Errorf("failled hash password")
	}

	input.Password = hasPass
	input.Role = &role
	return uc.userRepo.CreateUser(ctx, input)
}

func (uc *UseCase) CreateUserByAdmin(ctx context.Context, input user.RegisterByAdminRequest) error {
	tempPass := strings.Split(input.Email, "@")
	hasPass, err := helpers.HashPassword(tempPass[0])
	if err != nil {
		return fmt.Errorf("failled hash password")
	}

	input.Password = &hasPass
	return uc.userRepo.CreateUserByAdmin(ctx, input)
}

func (uc *UseCase) LoginUser(ctx context.Context, ipAddress string, input user.LoginRequest) (user.ResponseLogin, error) {
	data := user.ResponseLogin{}

	isExist := uc.logLoginRepo.CheckTableAlreadyExist(ctx)
	if !isExist {
		err := uc.logLoginRepo.CreateTableLogLogin(ctx)
		if err != nil {
			return data, err
		}
	}

	// cek apakah sudah login dan masih aktif tapi coba login kembali
	userAlreadyLogin, err := uc.logLoginRepo.GetLastLogLoginByEmail(ctx, input.Email)
	if userAlreadyLogin != nil && userAlreadyLogin.LogoutTime == nil {
		if time.Since(userAlreadyLogin.LoginTime) <= 2*time.Hour {
			err := uc.redisRepo.DeleteTokenRedis(input.Email)
			if err != nil {
				return data, fmt.Errorf("internal server error : %w", err)
			}

			now := time.Now()
			logLoginUpdate := loglogin.LogloginRequest{
				Email:       userAlreadyLogin.Email,
				FullName:    userAlreadyLogin.FullName,
				Role:        userAlreadyLogin.Role,
				IPAddress:   userAlreadyLogin.IPAddress,
				LoginTime:   userAlreadyLogin.LoginTime,
				LogoutTime:  &now,
				ProcessTime: &now,
				ExpiredTime: userAlreadyLogin.ExpiredTime,
			}

			err = uc.logLoginRepo.UpdateLogLogin(ctx, logLoginUpdate)
			if err != nil {
				return data, fmt.Errorf("internal server error : %w", err)
			}

			return data, fmt.Errorf("user is already logged in")
		}
	}

	userData, err := uc.userRepo.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return data, err
	}

	isValid := helpers.CheckPasswordHash(input.Password, userData.Password)
	if !isValid {
		return data, fmt.Errorf("invalid password/email")

	}

	dataToken := user.DataToken{
		ID:       userData.ID,
		FullName: userData.FullName,
		Role:     userData.Role,
		Email:    userData.Email,
	}

	token, err := helpers.CreateToken(dataToken)
	if err != nil {
		return data, fmt.Errorf("failled to generated token")

	}

	data = user.ResponseLogin{
		ID:       userData.ID,
		FullName: userData.FullName,
		Email:    userData.Email,
		Role:     userData.Role,
		Token:    token,
	}

	loginResponseJSON, err := json.Marshal(data)
	if err != nil {
		return data, fmt.Errorf("Internal server error : %w", err)
	}

	loginResponseString := string(loginResponseJSON)
	err = uc.redisRepo.SaveTokenInRedis(input.Email, loginResponseString)
	if err != nil {
		return data, fmt.Errorf("Internal server error : %w", err)
	}

	logLoginReq := loglogin.LogloginRequest{
		Email:     userData.Email,
		FullName:  userData.FullName,
		Role:      userData.Role,
		IPAddress: ipAddress,
	}

	err = uc.logLoginRepo.CreateLogLogin(ctx, logLoginReq)

	return data, nil
}

func (uc *UseCase) LogoutUser(ctx context.Context, email string) error {
	err := uc.redisRepo.DeleteTokenRedis(email)
	if err != nil {
		return err
	}

	userlogin, err := uc.logLoginRepo.GetLastLogLoginByEmail(ctx, email)
	if err != nil {
		return err
	}

	now := time.Now()
	dataLogin := loglogin.LogloginRequest{
		Email:       email,
		FullName:    userlogin.FullName,
		Role:        userlogin.Role,
		IPAddress:   userlogin.IPAddress,
		LoginTime:   userlogin.LoginTime,
		LogoutTime:  &now,
		ProcessTime: userlogin.ProcessTime,
		ExpiredTime: userlogin.ExpiredTime,
	}

	err = uc.logLoginRepo.UpdateLogLogin(ctx, dataLogin)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) UpdateRoleUser(ctx context.Context, userId int, input user.UpdateRoleRequest) error {
	return uc.userRepo.UpdateRoleUser(ctx, input.ID, userId, input.Role)
}

func (uc *UseCase) UpdatePasswordUser(ctx context.Context, userId int, input user.UpdatePaswordeRequest) error {

	if input.Password != input.ConfirmPassword {
		return fmt.Errorf("password and confirm password does not match")
	}

	hasPass, err := helpers.HashPassword(input.Password)
	if err != nil {
		return fmt.Errorf("failled hash password")
	}

	return uc.userRepo.UpdatePasswordUser(ctx, userId, hasPass)
}

func (uc *UseCase) GetUsers(ctx context.Context, input entity.Pagination) ([]user.User, int64, error) {
	data, total, err := uc.userRepo.GetUsers(ctx, input)
	if err != nil {
		return nil, 0, err
	}

	return data, total, nil
}
