package impl

import (
	m "github.com/fkmhrk/OpenInvoice/v1/model"
	company "github.com/fkmhrk/OpenInvoice/v1/model/company/mysql"
	"github.com/fkmhrk/OpenInvoice/v1/model/db"
	env "github.com/fkmhrk/OpenInvoice/v1/model/env/mysql"
	"github.com/fkmhrk/OpenInvoice/v1/model/logger"
	session "github.com/fkmhrk/OpenInvoice/v1/model/session/mysql"
	user "github.com/fkmhrk/OpenInvoice/v1/model/user/mysql"
)

// NewModels creates instance
func NewModels(connection *db.Connection) *m.Models {
	logger := logger.New()
	return &m.Models{
		User:           user.New(connection),
		Session:        session.NewSessionDAO(connection),
		SessionRefresh: session.NewSessionRefreshDAO(connection),
		Company:        company.New(connection),
		Trading:        NewTradingDAO(connection, logger),
		Env:            env.New(connection),
		Seq:            NewSeqDAO(connection),
		Logger:         logger,
	}
}
