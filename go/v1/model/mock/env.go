package mock

import (
	"github.com/fkmhrk/OpenInvoice/v1/entity"
)

type EnvDAO struct {
	CreateResult  entity.Env
	GetResult     entity.Env
	GetListResult []*entity.Env
	SaveResult    error
	UpdateResult  entity.Env
	DeleteResult  entity.Env
}

func (d *EnvDAO) Create(key, value string) (entity.Env, error) {
	return d.CreateResult, nil
}

func (d *EnvDAO) Get(key string) (entity.Env, error) {
	return d.GetResult, nil
}

func (d *EnvDAO) GetList() ([]*entity.Env, error) {
	return d.GetListResult, nil
}

func (d *EnvDAO) Save(list []*entity.Env) error {
	return d.SaveResult
}

func (d *EnvDAO) Update(key, value string) (entity.Env, error) {
	return d.UpdateResult, nil
}

func (d *EnvDAO) Delete(key string) (entity.Env, error) {
	return d.DeleteResult, nil
}
