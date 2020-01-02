package rest

import (
	"net/http"

	m "github.com/fkmhrk/OpenInvoice/v1/model"
	s "github.com/fkmhrk/OpenInvoice/v1/service"
	"github.com/fkmhrk/OpenInvoice/v1/service/company"
	"github.com/fkmhrk/OpenInvoice/v1/service/trading"
	"github.com/fkmhrk/OpenInvoice/v1/service/user"
	"github.com/gorilla/mux"
	"github.com/mokelab-go/hop"
)

const (
	method_GET    = "GET"
	method_POST   = "POST"
	method_PUT    = "PUT"
	method_DELETE = "DELETE"
)

func SetHandlers(r *mux.Router, services s.Services, u user.Service, t trading.Service, c company.Service, models *m.Models) {
	auth := hop.Operations(
		hop.GetPathParams,
		hop.GetCredential,
		getSession(models.Session),
	)
	authBody := hop.Operations(
		hop.GetPathParams,
		hop.GetCredential,
		getSession(models.Session),
		hop.GetBodyAsJSON,
	)

	r.HandleFunc("/token", hop.Operations(
		hop.GetContentType,
	)(getToken(u))).Methods(method_POST)
	r.HandleFunc("/token/refresh", hop.Operations(
		hop.GetBodyAsJSON,
	)(refreshToken(services))).Methods(method_POST)

	r.HandleFunc("/users",
		auth(getUsers(u))).
		Methods(method_GET)
	r.HandleFunc("/users",
		authBody(createUser(services))).
		Methods(method_POST)
	r.HandleFunc("/users/{id}",
		authBody(updateUser(services))).
		Methods(method_PUT)
	r.HandleFunc("/users/{id}",
		auth(deleteUser(services))).
		Methods(method_DELETE)
	r.HandleFunc("/tradings",
		auth(getTradings(t))).
		Methods(method_GET)
	r.HandleFunc("/tradings",
		authBody(createTrading(t))).
		Methods(method_POST)
	r.HandleFunc("/tradings/{tradingId}",
		auth(getTrading(t))).
		Methods(http.MethodGet)
	r.HandleFunc("/tradings/{tradingId}",
		authBody(updateTrading(t))).
		Methods(method_PUT)
	r.HandleFunc("/tradings/{tradingId}",
		auth(deleteTrading(services))).
		Methods(method_DELETE)
	r.HandleFunc("/tradings/{tradingId}/items",
		auth(getTradingItems(t))).
		Methods(method_GET)
	r.HandleFunc("/tradings/{tradingId}/items",
		authBody(createTradingItem(t))).
		Methods(method_POST)
	r.HandleFunc("/tradings/{tradingId}/items/{itemId}",
		authBody(updateTradingItem(t))).
		Methods(method_PUT)
	r.HandleFunc("/tradings/{tradingId}/items/{itemId}",
		auth(deleteTradingItem(t))).
		Methods(method_DELETE)
	r.HandleFunc("/companies",
		auth(getCompanies(c))).
		Methods(method_GET)
	r.HandleFunc("/companies",
		authBody(createCompany(c))).
		Methods(method_POST)
	r.HandleFunc("/companies/{companyId}",
		authBody(updateCompany(c))).
		Methods(method_PUT)
	r.HandleFunc("/companies/{companyId}",
		auth(deleteCompany(services))).
		Methods(method_DELETE)
	r.HandleFunc("/sequences/{seqType}",
		authBody(getNextNumber(services))).
		Methods(method_POST)
	// Environment
	r.HandleFunc("/environments",
		auth(getEnvironment(services))).
		Methods(method_GET)
	r.HandleFunc("/environments",
		authBody(saveEnvironment(services))).
		Methods(method_PUT)
	r.HandleFunc("/myCompany/name",
		getMyCompanyName(services)).
		Methods(method_GET)
}
