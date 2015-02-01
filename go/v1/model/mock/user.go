package mock

import (
	m "../"
)

type UserDAO struct {
	GetByNamePasswordResult *m.User
}

func (d *UserDAO) GetByNamePassword(name, password string) (*m.User, error) {
	return d.GetByNamePasswordResult, nil
}
