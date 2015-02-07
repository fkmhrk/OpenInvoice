package mock

import (
	m "../"
)

type UserDAO struct {
	GetByNamePasswordResult *m.User
	GetListResult           []*m.User
}

func (d *UserDAO) GetByNamePassword(name, password string) (*m.User, error) {
	return d.GetByNamePasswordResult, nil
}

func (d *UserDAO) GetList() ([]*m.User, error) {
	return d.GetListResult, nil
}
