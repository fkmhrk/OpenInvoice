package user

import (
	"github.com/fkmhrk/OpenInvoice/v1/model/session"
	"github.com/mokelab-go/server/entity"
)

type Service interface {
	GetToken(name, pass string) entity.Response
	RefreshToken(token string) entity.Response
	GetUsers() entity.Response
	Create(session *session.Session, loginName, displayName, tel, password string) entity.Response
	Update(session *session.Session, id, loginName, displayName, tel, password string) entity.Response
	Delete(session *session.Session, id string) entity.Response
}
