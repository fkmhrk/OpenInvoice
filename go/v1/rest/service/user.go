package service

import (
	e "github.com/fkmhrk/OpenInvoice/v1/entity"
	"github.com/mokelab-go/server/entity"
)

type User interface {
	GetToken(name, pass string) entity.Response
	RefreshToken(token string) entity.Response
	GetUsers() entity.Response
	Create(session *e.Session, loginName, displayName, tel, password string) entity.Response
	Update(session *e.Session, id, loginName, displayName, tel, password string) entity.Response
	Delete(session *e.Session, id string) entity.Response
}
