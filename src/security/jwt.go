package security

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	UserID    string `json:"user_id"`
	ExpiresAt int64  `json:"exp"`

	jwt.StandardClaims
}

var JWT_SECRET = []byte(os.Getenv("JWT_SECRET"))

func GenerateSignedJWT(userId string) (string, error) {
	claims := &JWTClaims{
		UserID:         userId,
		ExpiresAt:      time.Now().Add(time.Minute * 15).Unix(),
		StandardClaims: jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(JWT_SECRET)
}

func ValidateJWT(tokenString string) (*JWTClaims, error) {
	claims := &JWTClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			secret := JWT_SECRET
			return secret, nil
		},
	)

	if err != nil {
		return nil, errors.New("Something went wrong while parsing the token.")
	}
	if !token.Valid {
		return nil, errors.New("Invalid Token.")
	}

	expirationTime := time.Unix(claims.ExpiresAt, 0)
	if expirationTime.Before(time.Now()) {
		return nil, errors.New("Token is expired.")
	}

	return claims, nil
}
