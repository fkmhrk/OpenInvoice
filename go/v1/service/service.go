package service

import (
	user "github.com/fkmhrk/OpenInvoice/v1/service/user"
)

type Services struct {
	Admin   AdminService
	User    user.Service
	Trading TradingService
	Company CompanyService
}
