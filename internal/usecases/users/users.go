package users

import (
	"Task-Service-For-Teams/internal/entity"
	"database/sql"
	"log/slog"

	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type Repo interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) (*sql.Row, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}
type UseCase struct {
	logger *slog.Logger
	repo   Repo
}

func New(repo Repo, logger *slog.Logger) *UseCase {
	l := logger.With("service", "users")
	return &UseCase{
		logger: l,
		repo:   repo,
	}
}
func (uc *UseCase) CreateUser(u entity.User) (entity.User, error) {
	var createdUser entity.User
	hashPassword := uc.hashPassword(u.Password)
	q, err := uc.repo.QueryRow(`insert into users(name, email, password) values ($1,$2,$3) RETURNING (id, name, email)`, u.Name, u.Email, hashPassword)
	if err != nil {
		uc.logger.Error("err", err)
		return entity.User{}, err
	}
	err = q.Scan(&createdUser)
	if err != nil {
		return entity.User{}, err
	}
	return createdUser, nil

}
func (uc *UseCase) GetUsers(offset, limit int) ([]entity.User, error) {
	var users []entity.User
	q, err := uc.repo.Query(`select (id, email, name, passwrd) from users limit $1 offset $2`, limit, offset)
	if err != nil {
		uc.logger.Error("err", err)
		return nil, err
	}
	for q.Next() {
		var u entity.User
		err := q.Scan(&u)
		if err != nil {
			uc.logger.Error("err", err)
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
func (uc *UseCase) GetOneUser(id int) (entity.User, error)                   {}
func (uc *UseCase) UpdateUser(id int, user entity.User) (entity.User, error) {}
func (uc *UseCase) DeleteUser(id int) error                                  {}
func (uc *UseCase) hashPassword(password string) string {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), viper.GetInt("app.salt"))
	if err != nil {
		uc.logger.Error("password hashing error", err)
	}
	return string(passwordBytes)

}
