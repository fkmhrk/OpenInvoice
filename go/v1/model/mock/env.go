package mock

import (
	"github.com/fkmhrk/OpenInvoice/v1/model/env"
)

type EnvDAO struct {
	CreateResult  env.Env
	GetResult     env.Env
	GetListResult []*env.Env
	SaveResult    error
	UpdateResult  env.Env
	DeleteResult  env.Env
}

func (d *EnvDAO) Create(key, value string) (env.Env, error) {
	return d.CreateResult, nil
}

func (d *EnvDAO) Get(key string) (env.Env, error) {
	return d.GetResult, nil
}

func (d *EnvDAO) GetList() ([]*env.Env, error) {
	return d.GetListResult, nil
}

func (d *EnvDAO) Save(list []*env.Env) error {
	return d.SaveResult
}

func (d *EnvDAO) Update(key, value string) (env.Env, error) {
	return d.UpdateResult, nil
}

func (d *EnvDAO) Delete(key string) (env.Env, error) {
	return d.DeleteResult, nil
}
