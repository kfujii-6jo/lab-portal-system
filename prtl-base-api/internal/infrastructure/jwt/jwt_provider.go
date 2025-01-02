package jwt

import (
	"errors"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTProvider struct {
	privateKey string
	publicKey  string
}

func NewJWTProvider(privateKey string, publicKey string) *JWTProvider {
	return &JWTProvider{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

func (j *JWTProvider) GenerateToken(userID int) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": userID,
		"iss": "prtl-base-api",
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	slog.Info(j.privateKey)
	parsedPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(j.privateKey))
	tokenString, err := token.SignedString(parsedPrivateKey)
	return tokenString, err
}

func (j *JWTProvider) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.publicKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("Invalid token")
	}
	return token, nil
}
