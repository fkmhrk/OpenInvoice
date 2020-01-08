package mock

import (
	"github.com/fkmhrk/OpenInvoice/v1/model/user"
)

type UserDAO struct {
	GetByNamePasswordResult *user.User
	GetListResult           []*user.User
	GetByIdResult           *user.User
	CreateResult            *user.User
	UpdateResult            *user.User
	DeleteResult            error
}

func (d *UserDAO) GetByNamePassword(name, password string) (*user.User, error) {
	return d.GetByNamePasswordResult, nil
}

func (d *UserDAO) GetList() ([]*user.User, error) {
	return d.GetListResult, nil
}

func (d *UserDAO) GetById(id string) (*user.User, error) {
	return d.GetByIdResult, nil
}

func (d *UserDAO) Create(loginName, displayName, role, tel, password string) (*user.User, error) {
	return d.CreateResult, nil
}

func (d *UserDAO) Update(id, loginName, displayName, role, tel, password string) (*user.User, error) {
	return d.UpdateResult, nil
}

func (d *UserDAO) Delete(id string) error {
	return d.DeleteResult
}
