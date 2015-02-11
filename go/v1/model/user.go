package model

type UserDAO interface {
	GetByNamePassword(name, password string) (*User, error)
	GetList() ([]*User, error)
	Create(loginName, displayName, role, password string) (*User, error)
}

type User struct {
	Id           string
	LoginName    string
	DisplayName  string
	Role         Role
	CreatedTime  int64
	ModifiedTime int64
}
