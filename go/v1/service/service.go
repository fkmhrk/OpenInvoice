package service

import (
	"github.com/fkmhrk/OpenInvoice/v1/service/admin"
	"github.com/fkmhrk/OpenInvoice/v1/service/company"
	"github.com/fkmhrk/OpenInvoice/v1/service/trading"
	"github.com/fkmhrk/OpenInvoice/v1/service/user"
)

type Services struct {
	Admin   admin.Service
	User    user.Service
	Trading trading.Service
	Company company.Service
}
