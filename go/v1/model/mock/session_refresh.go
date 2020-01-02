package mock

import (
	m "github.com/fkmhrk/OpenInvoice/v1/model"
)

type SessionRefreshDAO struct {
	CreateResult m.SessionRefresh
	GetResult    m.SessionRefresh
	UpdateResult m.SessionRefresh
	DeleteResult m.SessionRefresh
}

func (d *SessionRefreshDAO) Create(userId, role string) (m.SessionRefresh, error) {
	return d.CreateResult, nil
}

func (d *SessionRefreshDAO) Get(token string) (m.SessionRefresh, error) {
	return d.GetResult, nil
}

func (d *SessionRefreshDAO) Update(token, userId, role string) (m.SessionRefresh, error) {
	return d.UpdateResult, nil
}

func (d *SessionRefreshDAO) Delete(token string) (m.SessionRefresh, error) {
	return d.DeleteResult, nil
}
