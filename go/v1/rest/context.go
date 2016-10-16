package rest

import (
	m "../model"
	"context"
)

type contextKey string

const (
	keySession contextKey = "app.session"
)

func setSession(c context.Context, session *m.Session) context.Context {
	return context.WithValue(c, keySession, session)
}

func session(c context.Context) *m.Session {
	if v, ok := c.Value(keySession).(*m.Session); ok {
		return v
	}
	return nil
}
