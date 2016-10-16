package service

import (
	m "../model"
)

type UserService interface {
	GetToken(name, pass string) Result
	RefreshToken(token string) Result
	GetUsers() Result
	Create(session *m.Session, loginName, displayName, tel, password string) Result
	Update(session *m.Session, id, loginName, displayName, tel, password string) Result
	Delete(session *m.Session, id string) Result
}
