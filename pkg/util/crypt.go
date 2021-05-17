package util

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(text string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	return string(hash)
}

func CompareHash(hash string, str string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(str)) == nil
}
