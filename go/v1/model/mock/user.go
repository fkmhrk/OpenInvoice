package mock

import (
	m "../"
)

type UserDAO struct {
	GetByNamePasswordResult *m.User
	GetListResult           []*m.User
	CreateResult            *m.User
}

func (d *UserDAO) GetByNamePassword(name, password string) (*m.User, error) {
	return d.GetByNamePasswordResult, nil
}

func (d *UserDAO) GetList() ([]*m.User, error) {
	return d.GetListResult, nil
}

func (d *UserDAO) Create(loginName, displayName, role, tel, password string) (*m.User, error) {
	return d.CreateResult, nil
}
