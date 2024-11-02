package middlewares

import (
	"context"
	"encoding/json"
	"net/http"
	"siap_app/internal/app/entity"
	"siap_app/internal/app/entity/user"
	"siap_app/internal/app/helpers"
	"strings"

	"github.com/go-redis/redis/v8"
)

func AuthorizationMiddleware(client *redis.Client) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var loginResponse user.ResponseLogin
			tokenString := extractTokenFromHeader(r)
			if tokenString == "" {
				helpers.SendUnauthorizedResponse(w)
				return
			}

			tokenData, err := helpers.VerifyToken(tokenString)
			if err != nil {
				helpers.SendUnauthorizedResponse(w)
				return
			}

			ctxRedis := context.Background()
			dataRedis, err := client.Get(ctxRedis, tokenData.Email).Result()
			if err != nil {
				helpers.SendUnauthorizedResponse(w)
				return
			}

			err = json.Unmarshal([]byte(dataRedis), &loginResponse)
			if err != nil {
				helpers.SendError(w, http.StatusInternalServerError, "internal server error", err.Error())
				return
			}

			if tokenString != loginResponse.Token {
				helpers.SendUnauthorizedResponse(w)
				return
			}

			ctx := context.WithValue(r.Context(), entity.UserIDKey, tokenData.UserId)
			ctx = context.WithValue(ctx, entity.FullNameKey, tokenData.FullName)
			ctx = context.WithValue(ctx, entity.RoleKey, tokenData.Role)
			ctx = context.WithValue(ctx, entity.EmailKey, tokenData.Email)
			ctx = context.WithValue(ctx, entity.IsAuthorizedKey, tokenData.IsAuthorized)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, ok := r.Context().Value(entity.RoleKey).(string)
		if !ok || role != "ADMIN" {
			helpers.SendForbiddenResponse(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func extractTokenFromHeader(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) == 2 && strings.EqualFold(parts[0], "Bearer") {
			return parts[1]
		}
	}
	return ""
}
