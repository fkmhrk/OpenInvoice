package v1

import (
	"github.com/fkmhrk/OpenInvoice/v1/model"
	"github.com/fkmhrk/OpenInvoice/v1/service"
	s "github.com/fkmhrk/OpenInvoice/v1/service"
	"github.com/fkmhrk/OpenInvoice/v1/service/admin"
	"github.com/fkmhrk/OpenInvoice/v1/service/company"
	"github.com/fkmhrk/OpenInvoice/v1/service/trading"
	"github.com/fkmhrk/OpenInvoice/v1/service/user"
)

func newServices(models *model.Models) s.Services {
	return service.Services{
		Admin:   admin.New(models),
		User:    user.New(models),
		Trading: trading.New(models),
		Company: company.New(models),
	}
}
