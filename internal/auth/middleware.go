package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func Validate(next http.Handler) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		tokenString := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer")

		claims := new(Claims)
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil {

			if err == jwt.ErrSignatureInvalid {
				http.Error(rw, err.Error(), http.StatusUnauthorized)
				return
			}
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		if !token.Valid {
			http.Error(rw, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", claims.UserId)
		ctx = context.WithValue(ctx, "username", claims.Username)
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}
