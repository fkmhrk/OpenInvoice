package rest

import (
	m "../model"
	s "../service"
	"github.com/gorilla/mux"
	"github.com/mokelab-go/hop"
	"net/http"
)

const (
	method_GET    = "GET"
	method_POST   = "POST"
	method_PUT    = "PUT"
	method_DELETE = "DELETE"
)

type handler func(http.ResponseWriter, *http.Request)

func SetHandlers(r *mux.Router, services s.Services, u s.UserService, t s.TradingService, c s.CompanyService, models *m.Models) {
	auth := hop.Operations(
		hop.GetPathParams,
		hop.GetCredential,
		getSession(models.Session),
	)

	r.HandleFunc("/token", hop.Operations(
		hop.GetContentType,
	)(getToken(u))).Methods(method_POST)
	r.HandleFunc("/token/refresh", hop.Operations(
		hop.GetBodyAsJSON,
	)(refreshToken(services))).Methods(method_POST)
	r.HandleFunc("/users", auth(getUsers(u))).
		Methods(method_GET)
	r.HandleFunc("/users", createUser(services)).
		Methods(method_POST)
	r.HandleFunc("/users/{id}", updateUser(services)).
		Methods(method_PUT)
	r.HandleFunc("/users/{id}", deleteUser(services)).
		Methods(method_DELETE)
	r.HandleFunc("/tradings", auth(getTradings(t))).
		Methods(method_GET)
	r.HandleFunc("/tradings", createTrading(t)).
		Methods(method_POST)
	r.HandleFunc("/tradings/{tradingId}", updateTrading(t)).
		Methods(method_PUT)
	r.HandleFunc("/tradings/{tradingId}", auth(deleteTrading(services))).
		Methods(method_DELETE)
	r.HandleFunc("/tradings/{tradingId}/items", auth(getTradingItems(t))).
		Methods(method_GET)
	r.HandleFunc("/tradings/{tradingId}/items", createTradingItem(t)).
		Methods(method_POST)
	r.HandleFunc("/tradings/{tradingId}/items/{itemId}", updateTradingItem(t)).
		Methods(method_PUT)
	r.HandleFunc("/tradings/{tradingId}/items/{itemId}",
		auth(deleteTradingItem(t))).
		Methods(method_DELETE)
	r.HandleFunc("/companies", getCompanies(c)).
		Methods(method_GET)
	r.HandleFunc("/companies", createCompany(c)).
		Methods(method_POST)
	r.HandleFunc("/companies/{companyId}", updateCompany(c)).
		Methods(method_PUT)
	r.HandleFunc("/companies/{companyId}", deleteCompany(services)).
		Methods(method_DELETE)
	r.HandleFunc("/sequences/{seqType}", getNextNumber(services)).
		Methods(method_POST)
	// Environment
	r.HandleFunc("/environments", getEnvironment(services)).
		Methods(method_GET)
	r.HandleFunc("/environments", saveEnvironment(services)).
		Methods(method_PUT)
	r.HandleFunc("/myCompany/name", getMyCompanyName(services)).
		Methods(method_GET)
}
