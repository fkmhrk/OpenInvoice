package model

type SessionRefreshDAO interface {
	Create(userId, role string) (SessionRefresh, error)
	Get(token string) (SessionRefresh, error)
	Update(token, userId, role string) (SessionRefresh, error)
	Delete(token string) (SessionRefresh, error)
}

type SessionRefresh struct {
	Token      string
	UserId     string
	Role       string
	ExpireTime int64
}

func (o SessionRefresh) IsEmpty() bool {
	return len(o.Token) == 0
}
