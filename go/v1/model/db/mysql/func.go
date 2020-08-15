package mysql

import (
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func generateUUID(length int) string {
	id1 := uuid.New().String()
	id2 := uuid.New().String()
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
	return "", errors.New("failed to insert 10 times")
}
