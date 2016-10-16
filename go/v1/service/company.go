package service

type CompanyService interface {
	GetList() Result
	Create(name, zip, address, phone, unit string) Result
	Update(id, name, zip, address, phone, unit string) Result
	Delete(id string) Result
}
