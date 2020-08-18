package entity

type SessionRefresh struct {
	Token      string
	UserId     string
	Role       Role
	ExpireTime int64
}

func (o *SessionRefresh) IsEmpty() bool {
	return len(o.Token) == 0
}
