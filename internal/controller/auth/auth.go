package auth

import (
	"log/slog"
	"net/http"
)

type AuthService interface {
}

type Auth struct {
	logger  *slog.Logger
	service *AuthService
}

func Register(w http.ResponseWriter, r *http.Request) {

}
func Login(w http.ResponseWriter, r *http.Request) {

}
