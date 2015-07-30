package service

type UserService interface {
	GetToken(name, pass string) Result
	RefreshToken(token string) Result
	GetUsers(token string) Result
	Create(token, loginName, displayName, tel, password string) Result
	Update(token, id, loginName, displayName, tel, password string) Result
	Delete(token, id string) Result
}
