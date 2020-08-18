package model

import "github.com/fkmhrk/OpenInvoice/v1/entity"

type Session interface {
	// Gets session by token
	GetByToken(token string) (*entity.Session, error)

	// Creates new session
	Create(userId, scope string, expireIn int64) (*entity.Session, error)
}
