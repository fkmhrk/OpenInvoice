package model

type UserDAO interface {
	GetByNamePassword(name, password string) (*User, error)
}

type User struct {
	Id           string
	LoginName    string
	CreatedTime  int64
	ModifiedTime int64
}
