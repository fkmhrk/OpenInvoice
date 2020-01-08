package v1

import (
	m "github.com/fkmhrk/OpenInvoice/v1/model"
	modeldb "github.com/fkmhrk/OpenInvoice/v1/model/db"
	mi "github.com/fkmhrk/OpenInvoice/v1/model/impl"
	"github.com/fkmhrk/OpenInvoice/v1/rest"
	s "github.com/fkmhrk/OpenInvoice/v1/service"
	si "github.com/fkmhrk/OpenInvoice/v1/service/impl"
	user "github.com/fkmhrk/OpenInvoice/v1/service/user"
	userImpl "github.com/fkmhrk/OpenInvoice/v1/service/user/impl"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func InitRouter(r *mux.Router) error {
	db, err := sql.Open("mysql", "openinvoice_test:openinvoice@tcp(db:3306)/openinvoice")
	if err != nil {
		return err
	}
	c := modeldb.NewConnection(db)
	models := mi.NewModels(c)
	services := si.NewServices(models)

	userService := userImpl.New(models.User, models.Session, models)
	initRouter(r, services, userService, models)
	return nil
}

func initRouter(r *mux.Router, services s.Services, u user.Service, models *m.Models) {
	r1 := r.PathPrefix("/api/v1").Subrouter()
	rest.SetHandlers(r1, services, u, services.Trading, services.Company, models)
}
