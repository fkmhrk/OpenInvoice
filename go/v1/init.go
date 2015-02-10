package v1

import (
	mi "./model/impl"
	"./rest"
	s "./service"
	si "./service/impl"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func InitRouter(r *mux.Router) error {
	db, err := sql.Open("mysql", "openinvoice_test:openinvoice@/openinvoice")
	if err != nil {
		return err
	}
	c := mi.NewConnection(db)
	userDAO := mi.NewUserDAO(c)
	sessionDAO := mi.NewSessionDAO(c)
	tradingDAO := mi.NewTradingDAO(c)
	companyDAO := mi.NewCompanyDAO(c)

	userService := si.NewUserSerivce(userDAO, sessionDAO)
	tradingService := si.NewTradingSerivce(sessionDAO, tradingDAO)
	companyService := si.NewCompanySerivce(sessionDAO, companyDAO)
	initRouter(r, userService, tradingService, companyService)
	return nil
}

func initRouter(r *mux.Router, u s.UserService, t s.TradingService, c s.CompanyService) {
	r1 := r.PathPrefix("/api/v1").Subrouter()
	rest.SetHandlers(r1, u, t, c)
}
