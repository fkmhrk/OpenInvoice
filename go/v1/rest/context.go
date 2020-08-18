package rest

import (
	"context"

	"github.com/fkmhrk/OpenInvoice/v1/entity"
)

type contextKey string

const (
	keySession contextKey = "app.session"
)

func setSession(c context.Context, session *entity.Session) context.Context {
	return context.WithValue(c, keySession, session)
}

func session(c context.Context) *entity.Session {
	if v, ok := c.Value(keySession).(*entity.Session); ok {
		return v
	}
	return nil
}
