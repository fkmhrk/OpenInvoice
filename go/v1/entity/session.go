package entity

// Session entity.
type Session struct {
	Token       string
	UserId      string
	Role        Role
	CreatedTime int64
	ExpireTime  int64
}
