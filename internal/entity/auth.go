package entity

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	jwt.RegisteredClaims
	Email string
	Name  string
	Teams []TeamInToken
}

type TeamInToken struct {
	Team *Teams
	Role string
}
