package middleware

import (
	"ecommerceapi/lib"

	"github.com/labstack/echo/v4/middleware"
)

var JWT = middleware.JWTWithConfig(middleware.JWTConfig{
	TokenLookup:    "cookie:" + lib.TokenCookieName,
	ParseTokenFunc: lib.ParseToken,
	SigningMethod:  lib.SigningMethod.Alg(),
})
