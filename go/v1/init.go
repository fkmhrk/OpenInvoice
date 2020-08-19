package v1

import (
	mi "github.com/fkmhrk/OpenInvoice/v1/model"
	modeldb "github.com/fkmhrk/OpenInvoice/v1/model/db"
	"github.com/fkmhrk/OpenInvoice/v1/rest"
	s "github.com/fkmhrk/OpenInvoice/v1/service"
	"github.com/fkmhrk/OpenInvoice/v1/service/model"

	"database/sql"

	// used for driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func InitRouter(r *mux.Router) error {
	db, err := sql.Open("mysql", "openinvoice_test:openinvoice@tcp(db:3306)/openinvoice")
	if err != nil {
		return err
	}
	c := modeldb.NewConnection(db)
	models := mi.New(c)
	services := newServices(models)

	initRouter(r, services, models)
	return nil
}

func initRouter(r *mux.Router, services s.Services, models *model.Models) {
	r1 := r.PathPrefix("/api/v1").Subrouter()
	rest.SetHandlers(r1, services, services.User, services.Trading, services.Company, models)
}
