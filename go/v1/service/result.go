package service

type Result interface {
	Status() int
	Body() string
	Headers() map[string]string
}
