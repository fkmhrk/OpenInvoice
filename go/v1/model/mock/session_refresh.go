package mock

import (
	"github.com/fkmhrk/OpenInvoice/v1/entity"
)

type SessionRefreshDAO struct {
	CreateResult entity.SessionRefresh
	GetResult    entity.SessionRefresh
	UpdateResult entity.SessionRefresh
	DeleteResult entity.SessionRefresh
}

func (d *SessionRefreshDAO) Create(userId, role string) (entity.SessionRefresh, error) {
	return d.CreateResult, nil
}

func (d *SessionRefreshDAO) Get(token string) (entity.SessionRefresh, error) {
	return d.GetResult, nil
}

func (d *SessionRefreshDAO) Update(token, userId, role string) (entity.SessionRefresh, error) {
	return d.UpdateResult, nil
}

func (d *SessionRefreshDAO) Delete(token string) (entity.SessionRefresh, error) {
	return d.DeleteResult, nil
}
