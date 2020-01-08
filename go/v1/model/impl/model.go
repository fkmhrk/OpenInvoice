package impl

import (
	m "github.com/fkmhrk/OpenInvoice/v1/model"
	company "github.com/fkmhrk/OpenInvoice/v1/model/company/mysql"
	"github.com/fkmhrk/OpenInvoice/v1/model/db"
	env "github.com/fkmhrk/OpenInvoice/v1/model/env/mysql"
	"github.com/fkmhrk/OpenInvoice/v1/model/logger"
	seq "github.com/fkmhrk/OpenInvoice/v1/model/seq/mysql"
	session "github.com/fkmhrk/OpenInvoice/v1/model/session/mysql"
	trading "github.com/fkmhrk/OpenInvoice/v1/model/trading/mysql"
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
		Trading:        trading.New(connection, logger),
		Env:            env.New(connection),
		Seq:            seq.New(connection),
		Logger:         logger,
	}
}
