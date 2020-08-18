package mock

import (
	"github.com/fkmhrk/OpenInvoice/v1/entity"
)

type CompanyDAO struct {
	GetListResult []*entity.Company
	GetByIdResult *entity.Company
	CreateResult  *entity.Company
	UpdateResult  *entity.Company
	DeleteResult  error
}

func (d *CompanyDAO) GetList() ([]*entity.Company, error) {
	return d.GetListResult, nil
}

func (d *CompanyDAO) GetById(id string) (*entity.Company, error) {
	return d.GetByIdResult, nil
}

func (d *CompanyDAO) Create(name, zip, address, phone, unit string) (*entity.Company, error) {
	return d.CreateResult, nil
}

func (d *CompanyDAO) Update(id, name, zip, address, phone, unit string) (*entity.Company, error) {
	return d.UpdateResult, nil
}

func (d *CompanyDAO) Delete(id string) error {
	return d.DeleteResult
}
