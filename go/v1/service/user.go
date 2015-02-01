package service

type UserService interface {
	GetToken(name, pass string) Result
}
