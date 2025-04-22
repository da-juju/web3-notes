package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd string) string {
	encodePwd, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(encodePwd)
}
