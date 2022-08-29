package lib

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

var TokenCookieName = os.Getenv("JWT_COOKIE_NAME")
var SigningMethod = jwt.SigningMethodHS256
var signingKey = os.Getenv("Signature")

type Claims struct {
	jwt.RegisteredClaims
	UserID int64 `json:"user_id"`
}

func ParseToken(tokenString string, c echo.Context) (interface{}, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != SigningMethod.Alg() {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", token.Header["alg"])
		}

		return []byte(signingKey), nil
	})

	err = VerifyToken(token, claims, err)

	return token, err
}

func CreateTokenString(userID int64, expirationTime time.Time) (string, error) {
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(SigningMethod, claims)

	return token.SignedString([]byte(signingKey))
}

func GetUserID(tokenString string) (int64, error) {
	claims := &Claims{}

	_, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if err != nil {
		return claims.UserID, err
	}

	return claims.UserID, nil
}

func VerifyToken(token *jwt.Token, claims *Claims, err error) error {
	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return errors.New("token expired")
	}

	return nil
}

func MapTokenToCookie(tokenString string, expirationTime time.Time) *http.Cookie {
	return &http.Cookie{
		Name:     TokenCookieName,
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
	}
}
