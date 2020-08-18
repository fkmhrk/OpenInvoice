package model

import "github.com/fkmhrk/OpenInvoice/v1/entity"

type User interface {
	GetByNamePassword(name, password string) (*entity.User, error)
	GetList() ([]*entity.User, error)
	GetById(id string) (*entity.User, error)
	Create(loginName, displayName, role, tel, password string) (*entity.User, error)
	Update(id, loginName, displayName, role, tel, password string) (*entity.User, error)
	Delete(id string) error
}
