package middleware

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func AuthAdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		token := strings.Split(header, " ")[1]
		slog.Info("asdasd")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if checkRoleFromJwt(token) != "admin" || checkRoleFromJwt(token) != "owner" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)

	})
}

func checkRoleFromJwt(tokenString string) string {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(viper.Get("app.jwt.secret").(string)), nil
	})
	if err != nil {
		return ""
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["role"].(string)
	}
	return ""
}
