package impl

import (
	s "../"
	m "../../model"
)

func NewServices(models *m.Models) s.Services {
	return s.Services{
		Admin:   NewAdminService(models),
		User:    NewUserSerivce(models.User, models.Session, models),
		Trading: NewTradingSerivce(models.Session, models.Trading, models),
		Company: NewCompanySerivce(models),
	}
}
