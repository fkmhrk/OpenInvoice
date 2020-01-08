package model

import (
	"github.com/fkmhrk/OpenInvoice/v1/model/company"
	"github.com/fkmhrk/OpenInvoice/v1/model/env"
	"github.com/fkmhrk/OpenInvoice/v1/model/logger"
	"github.com/fkmhrk/OpenInvoice/v1/model/seq"
	"github.com/fkmhrk/OpenInvoice/v1/model/session"
	"github.com/fkmhrk/OpenInvoice/v1/model/user"
)

// Models is collection
type Models struct {
	User           user.DAO
	Session        session.SessionDAO
	SessionRefresh session.SessionRefreshDAO
	Company        company.DAO
	Trading        TradingDAO
	Env            env.DAO
	Seq            seq.DAO
	Logger         logger.Logger
}
