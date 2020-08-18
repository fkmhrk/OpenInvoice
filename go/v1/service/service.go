package service

import (
	"github.com/fkmhrk/OpenInvoice/v1/rest/service"
	"github.com/fkmhrk/OpenInvoice/v1/service/admin"
	"github.com/fkmhrk/OpenInvoice/v1/service/company"
	"github.com/fkmhrk/OpenInvoice/v1/service/model"
	"github.com/fkmhrk/OpenInvoice/v1/service/trading"
	"github.com/fkmhrk/OpenInvoice/v1/service/user"
)

// New creates service instances
func New(models *model.Models) service.Services {
	return service.Services{
		Admin:   admin.New(models),
		User:    user.New(models),
		Trading: trading.New(models),
		Company: company.New(models),
	}
}
