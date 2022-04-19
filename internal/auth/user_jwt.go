// Package auth - JWT Provider
package auth

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

type UserClaims struct {
	jwt.StandardClaims
	Account string
	Role    string
}

func CreateUserToken(claims UserClaims) (string, error) {
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	if err != nil {
		return token, err
	}

	return token, nil
}

func VerifyUserToken(r *http.Request) (UserClaims, error) {
	givenToken := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	tokenClaims, err := jwt.ParseWithClaims(givenToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})

	// validate the signatures and expiration
	if err != nil {
		if verr, ok := err.(*jwt.ValidationError); ok {
			if verr.Errors&jwt.ValidationErrorMalformed != 0 {
				return UserClaims{}, errors.New("token is malformed")
			} else if verr.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				return UserClaims{}, errors.New("token signature is invalid")
			} else if verr.Errors&jwt.ValidationErrorExpired != 0 {
				return UserClaims{}, errors.New("token has been expired")
			} else if verr.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return UserClaims{}, errors.New("token is not valid yet")
			} else {
				return UserClaims{}, errors.New("cannot handle this token")
			}
		}
	}

	// validate custom claims
	if claim, ok := tokenClaims.Claims.(*UserClaims); ok && tokenClaims.Valid {
		if claim.Account != "" && claim.Role != "" {
			return *claim, nil
		}
	}

	return UserClaims{}, errors.New("token payload is invalid")
}
