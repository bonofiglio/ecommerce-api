package middleware

import (
	"ecommerceapi/server/auth"

	"github.com/labstack/echo/v4/middleware"
)

var JWT = middleware.JWTWithConfig(middleware.JWTConfig{
	TokenLookup:    "cookie:" + auth.TokenCookieName,
	ParseTokenFunc: auth.ParseToken,
	SigningMethod:  auth.SigningMethod.Alg(),
})
