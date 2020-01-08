// Copyright 2014 Mokelab Inc.  All rights reserved.

package session

import "github.com/fkmhrk/OpenInvoice/v1/model/user"

// DAO for session table.
type SessionDAO interface {
	// Gets session by token
	GetByToken(token string) (*Session, error)

	// Creates new session
	Create(userId, scope string, expireIn int64) (*Session, error)
}

// Session entity.
type Session struct {
	Token       string
	UserId      string
	Role        user.Role
	CreatedTime int64
	ExpireTime  int64
}
