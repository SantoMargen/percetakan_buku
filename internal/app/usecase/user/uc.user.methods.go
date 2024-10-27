package user

import (
	"context"
	"fmt"
	"siap_app/internal/app/entity/user"
	"siap_app/internal/app/helpers"
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

func (uc *UseCase) LoginUser(ctx context.Context, input user.LoginRequest) (user.ResponseLogin, error) {
	data := user.ResponseLogin{}
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
		Role:     userData.Role,
		Token:    token,
	}

	return data, nil
}
