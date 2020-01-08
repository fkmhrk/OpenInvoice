package model

import (
	"github.com/fkmhrk/OpenInvoice/v1/model/session"
	"github.com/fkmhrk/OpenInvoice/v1/model/user"
)

type Models struct {
	User           user.DAO
	Session        session.SessionDAO
	SessionRefresh session.SessionRefreshDAO
	Company        CompanyDAO
	Trading        TradingDAO
	Env            EnvDAO
	Seq            SeqDAO
	Logger         Logger
}
