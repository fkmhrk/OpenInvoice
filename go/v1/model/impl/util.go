package impl

import (
	"code.google.com/p/go-uuid/uuid"
	"crypto/sha1"
	"fmt"
	"strings"
)

func hashPassword(password string) string {
	password = "prefix" + password + "suffix"
	hashed := fmt.Sprintf("%x", sha1.Sum([]byte(password)))
	return hashed[0:32]
}

func generateSessionId() string {
	id1 := uuid.New()
	id2 := uuid.New()
	id := strings.Replace(id1+id2, "-", "", -1)
	return id[:48]
}

func generateId(length int) string {
	id1 := uuid.New()
	id2 := uuid.New()
	id := strings.Replace(id1+id2, "-", "", -1)
	return id[:length]
}
