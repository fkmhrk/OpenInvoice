package model

type EnvDAO interface {
	Create(key, value string) (Env, error)
	Get(key string) (Env, error)
	GetList() ([]*Env, error)
	Update(key, value string) (Env, error)
	Delete(key string) (Env, error)
}

type Env struct {
	Key   string
	Value string
}

func (o Env) IsEmpty() bool {
	return len(o.Key) == 0
}
