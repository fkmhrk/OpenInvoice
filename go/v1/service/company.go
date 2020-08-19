package service

import "github.com/mokelab-go/server/entity"

type Company interface {
	GetList() entity.Response
	Create(name, zip, address, phone, unit string) entity.Response
	Update(id, name, zip, address, phone, unit string) entity.Response
	Delete(id string) entity.Response
}
