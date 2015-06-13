package service

type UserService interface {
	GetToken(name, pass string) Result
	RefreshToken(token string) Result
	GetUsers(token string) Result
	Create(token, loginName, displayName, role, password string) Result
}
