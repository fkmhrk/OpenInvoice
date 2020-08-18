package model

import "github.com/fkmhrk/OpenInvoice/v1/entity"

type Company interface {
	// Gets all company
	GetList() ([]*entity.Company, error)
	GetById(id string) (*entity.Company, error)
	Create(name, zip, address, phone, unit string) (*entity.Company, error)
	Update(id, name, zip, address, phone, unit string) (*entity.Company, error)
	Delete(id string) error
}
