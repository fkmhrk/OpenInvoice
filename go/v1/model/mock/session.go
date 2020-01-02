package mock

import (
	m "github.com/fkmhrk/OpenInvoice/v1/model"
)

type SessionDAO struct {
	GetByTokenResult *m.Session
	CreateResult     *m.Session
}

func (d *SessionDAO) GetByToken(token string) (*m.Session, error) {
	return d.GetByTokenResult, nil
}

func (d *SessionDAO) Create(userId, scope string, expireIn int64) (*m.Session, error) {
	return d.CreateResult, nil
}
