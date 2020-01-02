package service

import (
	"github.com/fkmhrk/OpenInvoice/v1/service/company"
	user "github.com/fkmhrk/OpenInvoice/v1/service/user"
)

type Services struct {
	Admin   AdminService
	User    user.Service
	Trading TradingService
	Company company.Service
}
