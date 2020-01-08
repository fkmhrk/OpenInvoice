package mock

import (
	"github.com/fkmhrk/OpenInvoice/v1/model/session"
)

type SessionDAO struct {
	GetByTokenResult *session.Session
	CreateResult     *session.Session
}

func (d *SessionDAO) GetByToken(token string) (*session.Session, error) {
	return d.GetByTokenResult, nil
}

func (d *SessionDAO) Create(userId, scope string, expireIn int64) (*session.Session, error) {
	return d.CreateResult, nil
}
