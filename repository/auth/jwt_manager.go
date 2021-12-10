package auth

import (
	"fmt"
	"github.com/adarocket/controller/repository/user"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTManager -
type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
}

// NewJWTManager -
func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{secretKey, tokenDuration}
}

// UserClaims -
type UserClaims struct {
	jwt.StandardClaims
	Username    string   `json:"username"`
	Permissions []string `json:"permissions"`
	// Role        string   `json:"role"`
}

// Generate -
func (manager *JWTManager) Generate(user *user.User) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
		},
		Username:    user.Username,
		Permissions: user.Permissions,
		// Role:     user.Role,
	}

	/*
	   here I use a HMAC based signing method, which is HS256. For production, you should consider using stronger methods, Such as RSA or Eliptic-Curve based digital signature algorithms.
	*/
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.secretKey))
}

// Verify -
func (manager *JWTManager) Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				log.Println("unexpected token signing method")
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		log.Println("invalid token")
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		log.Println("invalid token claims")
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
