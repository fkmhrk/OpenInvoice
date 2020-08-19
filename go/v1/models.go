package v1

import (
	"github.com/fkmhrk/OpenInvoice/v1/model"
	"github.com/fkmhrk/OpenInvoice/v1/model/company"
	"github.com/fkmhrk/OpenInvoice/v1/model/db"
	"github.com/fkmhrk/OpenInvoice/v1/model/env"
	"github.com/fkmhrk/OpenInvoice/v1/model/logger"
	"github.com/fkmhrk/OpenInvoice/v1/model/seq"
	"github.com/fkmhrk/OpenInvoice/v1/model/session"
	"github.com/fkmhrk/OpenInvoice/v1/model/trading"
	"github.com/fkmhrk/OpenInvoice/v1/model/user"
)

func newModels(connection *db.Connection) *model.Models {
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
