package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID, secretKey string, ttl time.Duration) (string, error) {
	now := time.Now()
	claim := CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidationToken(tokenString, secretKey string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrTokenInvalidClaims
}
