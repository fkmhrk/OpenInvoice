package impl

import (
	m "github.com/fkmhrk/OpenInvoice/v1/model"
	s "github.com/fkmhrk/OpenInvoice/v1/service"
	company "github.com/fkmhrk/OpenInvoice/v1/service/company/impl"
	trading "github.com/fkmhrk/OpenInvoice/v1/service/trading/impl"
	user "github.com/fkmhrk/OpenInvoice/v1/service/user/impl"
)

func NewServices(models *m.Models) s.Services {
	return s.Services{
		Admin:   NewAdminService(models),
		User:    user.New(models.User, models.Session, models),
		Trading: trading.New(models.Session, models.Trading, models),
		Company: company.New(models),
	}
}
