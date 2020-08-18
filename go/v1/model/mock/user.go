package mock

import (
	"github.com/fkmhrk/OpenInvoice/v1/entity"
)

type UserDAO struct {
	GetByNamePasswordResult *entity.User
	GetListResult           []*entity.User
	GetByIdResult           *entity.User
	CreateResult            *entity.User
	UpdateResult            *entity.User
	DeleteResult            error
}

func (d *UserDAO) GetByNamePassword(name, password string) (*entity.User, error) {
	return d.GetByNamePasswordResult, nil
}

func (d *UserDAO) GetList() ([]*entity.User, error) {
	return d.GetListResult, nil
}

func (d *UserDAO) GetById(id string) (*entity.User, error) {
	return d.GetByIdResult, nil
}

func (d *UserDAO) Create(loginName, displayName, role, tel, password string) (*entity.User, error) {
	return d.CreateResult, nil
}

func (d *UserDAO) Update(id, loginName, displayName, role, tel, password string) (*entity.User, error) {
	return d.UpdateResult, nil
}

func (d *UserDAO) Delete(id string) error {
	return d.DeleteResult
}
