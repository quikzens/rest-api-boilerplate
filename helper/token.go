package helper

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/quikzens/rest-api-boilerplate/config"
)

// Different types of error returned by the verifyToken function
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type UserPayload struct {
	UserId    string    `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// Valid checks if the token payload is valid or not
func (payload *UserPayload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}

// CreateToken creates a new token with payload
func CreateToken(payload *UserPayload) (string, *UserPayload, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(config.TokenSecretKey))

	return token, payload, err
}

// VerifyToken checks if the token is valid or not
func VerifyToken(token string) (*UserPayload, error) {
	keyFunc := func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}

		return []byte(config.TokenSecretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &UserPayload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}

		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*UserPayload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
