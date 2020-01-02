package mock

import (
	m "github.com/fkmhrk/OpenInvoice/v1/model"
)

type EnvDAO struct {
	CreateResult  m.Env
	GetResult     m.Env
	GetListResult []*m.Env
	SaveResult    error
	UpdateResult  m.Env
	DeleteResult  m.Env
}

func (d *EnvDAO) Create(key, value string) (m.Env, error) {
	return d.CreateResult, nil
}

func (d *EnvDAO) Get(key string) (m.Env, error) {
	return d.GetResult, nil
}

func (d *EnvDAO) GetList() ([]*m.Env, error) {
	return d.GetListResult, nil
}

func (d *EnvDAO) Save(list []*m.Env) error {
	return d.SaveResult
}

func (d *EnvDAO) Update(key, value string) (m.Env, error) {
	return d.UpdateResult, nil
}

func (d *EnvDAO) Delete(key string) (m.Env, error) {
	return d.DeleteResult, nil
}
