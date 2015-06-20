package service

type Services struct {
	Admin   AdminService
	User    UserService
	Trading TradingService
	Company CompanyService
}
