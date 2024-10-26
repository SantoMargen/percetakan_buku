package helpers

import (
	"errors"
	"fmt"
	"os"
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

type TokenData struct {
	UserId       int
	FullName     string
	Role         string
	Email        string
	IsAuthorized bool
}

func VerifyToken(tokenString string) (TokenData, error) {
	claims := &jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return TokenData{}, fmt.Errorf("error parsing token: %w", err)
	}

	if !token.Valid {
		return TokenData{}, errors.New("invalid token")
	}

	userId, ok := (*claims)["userId"].(float64)
	if !ok {
		return TokenData{}, errors.New("userId claim is not a valid number")
	}

	isAuthorized, _ := (*claims)["authorized"].(bool)
	role, _ := (*claims)["role"].(string)
	fullName, _ := (*claims)["fullName"].(string)
	email, _ := (*claims)["email"].(string)

	return TokenData{
		UserId:       int(userId),
		IsAuthorized: isAuthorized,
		Role:         role,
		FullName:     fullName,
		Email:        email,
	}, nil
}
