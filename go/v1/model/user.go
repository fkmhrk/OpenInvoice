package model

type UserDAO interface {
	GetByNamePassword(name, password string) (*User, error)
	GetList() ([]*User, error)
	Create(loginName, displayName, role, tel, password string) (*User, error)
	Update(id, loginName, displayName, role, tel, password string) (*User, error)
}

type User struct {
	Id           string
	LoginName    string
	DisplayName  string
	Role         Role
	Tel          string
	CreatedTime  int64
	ModifiedTime int64
}
