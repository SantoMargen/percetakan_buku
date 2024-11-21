package helpers

import (
	"errors"
	"fmt"
	"os"
	"siap_app/internal/app/entity"
	"siap_app/internal/app/entity/user"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func CreateToken(input user.DataToken) (string, error) {
	Claims := jwt.MapClaims{}
	Claims["authorized"] = true
	Claims["userId"] = input.ID
	Claims["email"] = input.Email
	Claims["full_name"] = input.FullName
	Claims["role"] = input.Role
	Claims["exp"] = time.Now().Add(time.Hour * 10).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	return token.SignedString(secretKey)
}

func VerifyToken(tokenString string) (entity.TokenData, error) {
	var resp entity.TokenData
	claims := &jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return resp, fmt.Errorf("error parsing token: %w", err)
	}

	if !token.Valid {
		return resp, errors.New("invalid token")
	}

	userId, ok := (*claims)["userId"].(float64)
	if !ok {
		return resp, errors.New("userId claim is not a valid number")
	}

	isAuthorized, _ := (*claims)["authorized"].(bool)
	role, _ := (*claims)["role"].(string)
	fullName, _ := (*claims)["fullName"].(string)
	email, _ := (*claims)["email"].(string)

	resp = entity.TokenData{
		UserId:       int(userId),
		IsAuthorized: isAuthorized,
		Role:         role,
		FullName:     fullName,
		Email:        email,
	}
	return resp, nil
}
