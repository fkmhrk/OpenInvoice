package user

import (
	"crypto/sha1"
	"fmt"
)

func hashPassword(password string) string {
	password = "prefix" + password + "suffix"
	hashed := fmt.Sprintf("%x", sha1.Sum([]byte(password)))
	return hashed[0:32]
}
