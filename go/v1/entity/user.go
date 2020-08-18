package entity

type User struct {
	Id           string
	LoginName    string
	DisplayName  string
	Role         Role
	Tel          string
	CreatedTime  int64
	ModifiedTime int64
}
