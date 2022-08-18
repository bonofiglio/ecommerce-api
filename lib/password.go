package lib

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
)

func HashPassword(password string, salt string) string {
	passwordPepper := os.Getenv("PASSWORD_PEPPER")

	passwordString := password + passwordPepper + salt

	hash := sha256.Sum256([]byte(passwordString))

	return hex.EncodeToString(hash[:])
}
