package services

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v4"
)

type TokenService struct {
	tokenSecret []byte
}

func NewTokenService() *TokenService {
	return &TokenService{
		tokenSecret: []byte{},
	}
}

func (service *TokenService) ParseToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		log.Printf(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return service.tokenSecret, nil
}

func (service *TokenService) GetClaim(jwtToken string, claimKey string) (string, error) {
	token, err := jwt.Parse(jwtToken, service.ParseToken)
	if err != nil {
		return "", err
	}
	var value interface{}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		value = claims[claimKey]
		if value == nil {
			log.Printf("Error!! = > %s", value)
			return "", fmt.Errorf("Key %s not found", claimKey)
		}
	} else {
		return "", fmt.Errorf("Invalid token")
	}
	return value.(string), nil
}
