package mock

import (
	m "../"
)

func NewMock() *m.Models {
	return &m.Models{
		Session: &SessionDAO{},
		Env:     &EnvDAO{},
	}
}
