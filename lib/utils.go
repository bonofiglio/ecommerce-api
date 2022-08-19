package lib

import (
	"database/sql"
	"math/rand"
	"strings"

	"github.com/labstack/echo/v4"
)

var base64Characters = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

func GenerateRandomString(length int) string {
	sb := strings.Builder{}
	// Allocate enough memory for the string
	sb.Grow(length)

	// Generate a random string of "length" characters
	for i := 0; i < length; i++ {
		// Pick a random character from the base64Characters array and write it to the string builder
		sb.WriteByte(base64Characters[rand.Int63()%int64(len(base64Characters))])
	}

	return sb.String()
}

func PanicOnDbError(result sql.Result, err error) sql.Result {
	if err != nil {
		panic(err)
	}

	return result
}

func CreateNewResponseError(status int, messages ...string) *echo.HTTPError {
	return echo.NewHTTPError(status, map[string]interface{}{
		"errors": messages,
	})
}
