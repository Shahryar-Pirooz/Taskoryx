package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	secretKey  []byte
	accessTTL  time.Duration
	refreshTTL time.Duration
}

func NewJWTManager(secretKey []byte, accessTTL, refreshTTL time.Duration) *JWTManager {
	return &JWTManager{
		secretKey:  secretKey,
		accessTTL:  accessTTL,
		refreshTTL: refreshTTL,
	}
}

func (jm *JWTManager) generateToken(userID string, ttl time.Duration, now time.Time) (string, error) {
	claim := CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(jm.secretKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (jm *JWTManager) GenerateAccessToken(userID string) (string, error) {
	now := time.Now()
	return jm.generateToken(userID, jm.accessTTL, now)
}
func (jm *JWTManager) GenerateRefreshToken(userID string) (string, error) {
	now := time.Now()
	return jm.generateToken(userID, jm.refreshTTL, now)
}

func (jm *JWTManager) GeneratePairToken(userID string) (string, string, error) {
	accessToken, err := jm.GenerateAccessToken(userID)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := jm.GenerateRefreshToken(userID)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (jm *JWTManager) ValidationToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return jm.secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrTokenInvalidClaims
}
