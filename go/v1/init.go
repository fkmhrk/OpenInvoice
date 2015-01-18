package v1

import (
	"./rest"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func InitRouter(r *mux.Router) error {
	_, err := sql.Open("mysql", "openinvoice_test:openinvocie@/openinvoce")
	if err != nil {
		return err
	}

	initRouter(r)
	return nil
}

func initRouter(r *mux.Router) {
	r1 := r.PathPrefix("/api/v1").Subrouter()
	rest.SetHandlers(r1)
}
