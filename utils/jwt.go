package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecretKey []byte

// SetJWTSecretKey initializes the JWT secret key from an external source, such as environment variables.
func SetJWTSecretKey(secret string) {
	jwtSecretKey = []byte(secret)
}

// GenerateJWT generates a JWT token for the given user ID.
func GenerateJWT(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub":    userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		"issued": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

// ParseJWT validates a JWT token and retrieves the user ID.
func ParseJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token claims")
	}

	userID, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("userId not found in token")
	}

	return userID, nil
}
