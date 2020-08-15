package impl

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/pborman/uuid"
)

func generateUUID(length int) string {
	id1 := uuid.New()
	id2 := uuid.New()
	id := strings.Replace(id1+id2, "-", "", -1)
	return id[:length]
}

func InsertWithUUID(l int, f func(id string) error) (string, error) {
	for i := 0; i < 10; i++ {
		id := generateUUID(l)
		err := f(id)
		if err == nil {
			return id, nil
		}
		if err2, ok := err.(*mysql.MySQLError); ok {
			if err2.Number != 1062 {
				return "", err2
			}
		} else {
			return "", err
		}

	}
	return "", errors.New("Failed to insert 10 times.")
}

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
