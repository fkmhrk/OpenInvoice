package mock

import (
	"github.com/fkmhrk/OpenInvoice/v1/model/session"
)

type SessionRefreshDAO struct {
	CreateResult session.SessionRefresh
	GetResult    session.SessionRefresh
	UpdateResult session.SessionRefresh
	DeleteResult session.SessionRefresh
}

func (d *SessionRefreshDAO) Create(userId, role string) (session.SessionRefresh, error) {
	return d.CreateResult, nil
}

func (d *SessionRefreshDAO) Get(token string) (session.SessionRefresh, error) {
	return d.GetResult, nil
}

func (d *SessionRefreshDAO) Update(token, userId, role string) (session.SessionRefresh, error) {
	return d.UpdateResult, nil
}

func (d *SessionRefreshDAO) Delete(token string) (session.SessionRefresh, error) {
	return d.DeleteResult, nil
}
