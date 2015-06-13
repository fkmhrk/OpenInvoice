package model

type Models struct {
	User           UserDAO
	Session        SessionDAO
	SessionRefresh SessionRefreshDAO
	Company        CompanyDAO
	Trading        TradingDAO
	Env            EnvDAO
	Seq            SeqDAO
	Logger         Logger
}
