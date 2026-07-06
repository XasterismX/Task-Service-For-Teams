package auth

import (
	"Task-Service-For-Teams/internal/entity"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	logger      *slog.Logger
	teamService RoleFiner
}
type RoleFiner interface {
	GetTeamsByUserId(id int64) ([]entity.TeamInToken, error)
}

func NewAuthService(logger *slog.Logger) *AuthService {
	return &AuthService{logger: logger}
}
func (auc *AuthService) CreateJwtToken(user entity.User) (string, error) {
	teams, err := auc.teamService.GetTeamsByUserId(user.Id)
	if err != nil {
		return "", err
	}
	t := jwt.New(jwt.SigningMethodHS256)
	t.Claims = &entity.UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
		Email: user.Email,
		Name:  user.Name,
		Teams: teams,
	}
}
func (auc *AuthService) ValidateJwt(token string) (bool, error) {

}
func (auc *AuthService) GetClaimsFromJwt() (jwt.MapClaims, error) {}
