package rest

import (
	"context"

	ses "github.com/fkmhrk/OpenInvoice/v1/model/session"
)

type contextKey string

const (
	keySession contextKey = "app.session"
)

func setSession(c context.Context, session *ses.Session) context.Context {
	return context.WithValue(c, keySession, session)
}

func session(c context.Context) *ses.Session {
	if v, ok := c.Value(keySession).(*ses.Session); ok {
		return v
	}
	return nil
}
