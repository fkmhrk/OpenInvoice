package service

type UserService interface {
	GetToken(name, pass string) Result
	GetUsers(token string) Result
}
