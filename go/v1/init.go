package v1

import (
	m "./model"
	mi "./model/impl"
	"./rest"
	s "./service"
	si "./service/impl"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func InitRouter(r *mux.Router) error {
	db, err := sql.Open("mysql", "openinvoice_test:openinvoice@tcp(db:3306)/openinvoice")
	if err != nil {
		return err
	}
	c := mi.NewConnection(db)
	models := mi.NewModels(c)
	services := si.NewServices(models)

	userService := si.NewUserSerivce(models.User, models.Session, models)
	tradingService := si.NewTradingSerivce(models.Session, models.Trading, models)
	companyService := si.NewCompanySerivce(models)
	initRouter(r, services, userService, tradingService, companyService, models)
	return nil
}

func initRouter(r *mux.Router, services s.Services, u s.UserService, t s.TradingService, c s.CompanyService, models *m.Models) {
	r1 := r.PathPrefix("/api/v1").Subrouter()
	rest.SetHandlers(r1, services, u, t, c, models)
}
