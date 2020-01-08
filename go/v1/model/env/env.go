package env

// DAO is interface
type DAO interface {
	Create(key, value string) (Env, error)
	Get(key string) (Env, error)
	GetList() ([]*Env, error)
	Save(list []*Env) error
	Update(key, value string) (Env, error)
	Delete(key string) (Env, error)
}

// Env is
type Env struct {
	Key   string
	Value string
}

// IsEmpty determines this is empty
func (o Env) IsEmpty() bool {
	return len(o.Key) == 0
}
