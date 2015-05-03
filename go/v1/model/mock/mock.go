package mock

import (
	m "../"
)

func NewMock() *m.Models {
	return &m.Models{
		Env: &EnvDAO{},
	}
}
