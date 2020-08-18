package mock

import "github.com/fkmhrk/OpenInvoice/v1/entity"

type SessionDAO struct {
	GetByTokenResult *entity.Session
	CreateResult     *entity.Session
}

func (d *SessionDAO) GetByToken(token string) (*entity.Session, error) {
	return d.GetByTokenResult, nil
}

func (d *SessionDAO) Create(userId, scope string, expireIn int64) (*entity.Session, error) {
	return d.CreateResult, nil
}
