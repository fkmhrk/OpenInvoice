package model

import (
	"github.com/fkmhrk/OpenInvoice/v1/model/company"
	"github.com/fkmhrk/OpenInvoice/v1/model/logger"
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
	Env            EnvDAO
	Seq            SeqDAO
	Logger         logger.Logger
}
