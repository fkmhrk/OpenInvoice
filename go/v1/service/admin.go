package service

type AdminService interface {
	// Gets Environment
	GetEnvironment(token string) Result
}
