package model

type Models struct {
	User    UserDAO
	Session SessionDAO
	Company CompanyDAO
	Trading TradingDAO
	Env     EnvDAO
	Seq     SeqDAO
	Logger  Logger
}
