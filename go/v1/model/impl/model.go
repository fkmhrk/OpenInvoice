package impl

import (
	m "github.com/fkmhrk/OpenInvoice/v1/model"
	"github.com/fkmhrk/OpenInvoice/v1/model/db"
	user "github.com/fkmhrk/OpenInvoice/v1/model/user/mysql"
	session "github.com/fkmhrk/OpenInvoice/v1/model/session/mysql"
)

func NewModels(connection *db.Connection) *m.Models {
	logger := NewLogger()
	return &m.Models{
		User:           user.New(connection),
		Session:        session.NewSessionDAO(connection),
		SessionRefresh: session.NewSessionRefreshDAO(connection),
		Company:        NewCompanyDAO(connection),
		Trading:        NewTradingDAO(connection, logger),
		Env:            NewEnvDAO(connection),
		Seq:            NewSeqDAO(connection),
		Logger:         logger,
	}
}
