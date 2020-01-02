package user

import (
	m "github.com/fkmhrk/OpenInvoice/v1/model"
	"github.com/mokelab-go/server/entity"
)

type Service interface {
	GetToken(name, pass string) entity.Response
	RefreshToken(token string) entity.Response
	GetUsers() entity.Response
	Create(session *m.Session, loginName, displayName, tel, password string) entity.Response
	Update(session *m.Session, id, loginName, displayName, tel, password string) entity.Response
	Delete(session *m.Session, id string) entity.Response
}
