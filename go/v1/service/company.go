package service

type CompanyService interface {
	GetList(token string) Result
	Create(token, name, zip, address, phone, unit string) Result
	Update(token, id, name, zip, address, phone, unit string) Result
	Delete(token, id string) Result
}
