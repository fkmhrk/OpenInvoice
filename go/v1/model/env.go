package model

import "github.com/fkmhrk/OpenInvoice/v1/entity"

type Env interface {
	Create(key, value string) (entity.Env, error)
	Get(key string) (entity.Env, error)
	GetList() ([]*entity.Env, error)
	Save(list []*entity.Env) error
	Update(key, value string) (entity.Env, error)
	Delete(key string) (entity.Env, error)
}
