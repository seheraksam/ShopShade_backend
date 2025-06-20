package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID string, role string) (string, string, error) {
	jwtKey := os.Getenv("JWT_SECRET")
	if jwtKey == "" {
		return "", "", errors.New("JWT_SECRET is not set")
	}
	fmt.Println(jwtKey)
	accessClaims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}
	fmt.Println(accessClaims)
	refreshClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	}
	fmt.Println(accessClaims)
	accessT := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refreshT := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	fmt.Println(accessT)
	fmt.Println(refreshT)
	accessToken, err := accessT.SignedString(jwtKey)
	if err != nil {
		return "", "", errors.New("access token error")
	}
	fmt.Println(accessToken)

	refreshToken, err := refreshT.SignedString(jwtKey)
	if err != nil {
		return "", "", errors.New("refresh token error")
	}
	fmt.Println(refreshToken)
	return accessToken, refreshToken, nil
}
