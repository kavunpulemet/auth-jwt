package models

import "github.com/golang-jwt/jwt/v5"

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Claims struct {
	RefreshTokenID string `json:"refresh_token_id"`
	UserID         string `json:"user_id"`
	UserIP         string `json:"user_ip"`
	jwt.RegisteredClaims
}
