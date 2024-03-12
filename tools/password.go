package tools

import (
	"encoding/hex"

	bcrypt "golang.org/x/crypto/bcrypt"
	sh3a "golang.org/x/crypto/sha3"
)

func HashPassword(pass string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPass), nil
}

func CheckPassword(pass string, hashedPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
}

func Hash256Password(pass string) string {
	buf := []byte(pass)
	pwd := sh3a.New256()
	pwd.Write(buf)

	return hex.EncodeToString(pwd.Sum(nil))
}
