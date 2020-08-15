package impl

import (
	m "github.com/fkmhrk/OpenInvoice/v1/model"
	"github.com/fkmhrk/OpenInvoice/v1/rest/service"
	admin "github.com/fkmhrk/OpenInvoice/v1/service/admin/impl"
	company "github.com/fkmhrk/OpenInvoice/v1/service/company/impl"
	trading "github.com/fkmhrk/OpenInvoice/v1/service/trading/impl"
	user "github.com/fkmhrk/OpenInvoice/v1/service/user/impl"
)

func NewServices(models *m.Models) service.Services {
	return service.Services{
		Admin:   admin.New(models),
		User:    user.New(models.User, models.Session, models),
		Trading: trading.New(models.Session, models.Trading, models),
		Company: company.New(models),
	}
}
