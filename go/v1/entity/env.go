package entity

// Env is
type Env struct {
	Key   string
	Value string
}

// IsEmpty determines this is empty
func (o Env) IsEmpty() bool {
	return len(o.Key) == 0
}
