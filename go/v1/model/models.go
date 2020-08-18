package model

import (
	company "github.com/fkmhrk/OpenInvoice/v1/model/company/mysql"
	"github.com/fkmhrk/OpenInvoice/v1/model/db"
	env "github.com/fkmhrk/OpenInvoice/v1/model/env/mysql"
	"github.com/fkmhrk/OpenInvoice/v1/model/logger"
	seq "github.com/fkmhrk/OpenInvoice/v1/model/seq/mysql"
	session "github.com/fkmhrk/OpenInvoice/v1/model/session/mysql"
	trading "github.com/fkmhrk/OpenInvoice/v1/model/trading/mysql"
	user "github.com/fkmhrk/OpenInvoice/v1/model/user/mysql"
	"github.com/fkmhrk/OpenInvoice/v1/service/model"
)

// New creates model instances
func New(connection *db.Connection) *model.Models {
	logger := logger.New()
	return &model.Models{
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
