package models

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	UserID uint   `json:"userId"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}
