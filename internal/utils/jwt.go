/*
Date Created		1 May 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Website				https://mikeio.web.app
Status				Looking for a job!
Description			A Fintech Data Service
Inspired by			https://freecurrencyapi.com
*/

package utils

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mikietechie/gocurrenciesapi/internal/config"
)

func GerateToken(str string) (string, error) {

	// Create a new JWT token with claims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": str,                                              // Subject (user identifier)
		"iss": config.SYS_NAME,                                  // Issuer
		"exp": time.Now().Add(config.JWT_TOKEN_LIFETIME).Unix(), // Expiration time
		"iat": time.Now().Unix(),                                // Issued at
	})
	tokenString, err := claims.SignedString([]byte(config.SECRET_KEY))
	if err != nil {
		return "", err
	}
	// Print information about the created token
	log.Println("Token claims added\t:\n", claims)
	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	// Parse the token with the secret key
	var token *jwt.Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SECRET_KEY), nil
	})

	// Check for verification errors
	if err != nil {
		return token, err
	}

	// Check if the token is valid
	if !token.Valid {
		return token, errors.New("invalid token")
	}

	// Return the verified token
	return token, nil
}
