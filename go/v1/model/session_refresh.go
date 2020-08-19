package model

import "github.com/fkmhrk/OpenInvoice/v1/entity"

type SessionRefresh interface {
	Create(userId, role string) (entity.SessionRefresh, error)
	Get(token string) (entity.SessionRefresh, error)
	Update(token, userId, role string) (entity.SessionRefresh, error)
	Delete(token string) (entity.SessionRefresh, error)
}
