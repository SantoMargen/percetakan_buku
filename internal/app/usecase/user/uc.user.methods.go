package user

import (
	"context"
	"encoding/json"
	"fmt"
	loglogin "siap_app/internal/app/entity/log_login"
	"siap_app/internal/app/entity/user"
	"siap_app/internal/app/helpers"
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
	hasPass, err := helpers.HashPassword(input.Password)
	if err != nil {
		return fmt.Errorf("failled hash password")
	}

	input.Password = hasPass
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
	fmt.Println(err, "=============================== MASUK")
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
