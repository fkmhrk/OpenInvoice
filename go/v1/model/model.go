package model

type Models struct {
	User    UserDAO
	Session SessionDAO
	Company CompanyDAO
	Trading TradingDAO
	Env     EnvDAO
	Logger  Logger
}
