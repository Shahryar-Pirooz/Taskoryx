package jwt

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	UserID string `json:"user_id"`
	Role   int8   `json:"role"`
	jwt.RegisteredClaims
}
