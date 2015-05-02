package model

type Models struct {
	User    UserDAO
	Session SessionDAO
	Company CompanyDAO
	Trading TradingDAO
	Logger  Logger
}
