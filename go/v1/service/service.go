package service

import (
	"github.com/fkmhrk/OpenInvoice/v1/service/company"
	"github.com/fkmhrk/OpenInvoice/v1/service/trading"
	user "github.com/fkmhrk/OpenInvoice/v1/service/user"
)

type Services struct {
	Admin   AdminService
	User    user.Service
	Trading trading.Service
	Company company.Service
}
