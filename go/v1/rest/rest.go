package rest

import (
	"net/http"

	"github.com/fkmhrk/OpenInvoice/v1/model"
	"github.com/fkmhrk/OpenInvoice/v1/service"
	"github.com/gorilla/mux"
	"github.com/mokelab-go/hop"
)

const (
	method_GET    = "GET"
	method_POST   = "POST"
	method_PUT    = "PUT"
	method_DELETE = "DELETE"
)

func SetHandlers(r *mux.Router, services service.Services, u service.User, t service.Trading, c service.Company, models *model.Models) {
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
		auth(getTradings(services.Trading))).
		Methods(method_GET)
	r.HandleFunc("/tradings",
		authBody(createTrading(services.Trading))).
		Methods(method_POST)
	r.HandleFunc("/tradings/{tradingId}",
		auth(getTrading(services.Trading))).
		Methods(http.MethodGet)
	r.HandleFunc("/tradings/{tradingId}",
		authBody(updateTrading(services.Trading))).
		Methods(method_PUT)
	r.HandleFunc("/tradings/{tradingId}",
		auth(deleteTrading(services))).
		Methods(method_DELETE)
	r.HandleFunc("/tradings/{tradingId}/items",
		auth(getTradingItems(services.Trading))).
		Methods(method_GET)
	r.HandleFunc("/tradings/{tradingId}/items",
		authBody(createTradingItem(services.Trading))).
		Methods(method_POST)
	r.HandleFunc("/tradings/{tradingId}/items/{itemId}",
		authBody(updateTradingItem(services.Trading))).
		Methods(method_PUT)
	r.HandleFunc("/tradings/{tradingId}/items/{itemId}",
		auth(deleteTradingItem(services.Trading))).
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
