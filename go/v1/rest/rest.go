package rest

import (
	s "../service"
	"github.com/gorilla/mux"
	"net/http"
)

type handler func(http.ResponseWriter, *http.Request)

func SetHandlers(r *mux.Router, services s.Services, u s.UserService, t s.TradingService, c s.CompanyService) {
	r.HandleFunc("/token", getToken(u)).
		Methods("POST")
	r.HandleFunc("/users", getUsers(u)).
		Methods("GET")
	r.HandleFunc("/tradings", getTradings(t)).
		Methods("GET")
	r.HandleFunc("/tradings", createTrading(t)).
		Methods("POST")
	r.HandleFunc("/tradings/{tradingId}", updateTrading(t)).
		Methods("PUT")
	r.HandleFunc("/tradings/{tradingId}/items", getTradingItems(t)).
		Methods("GET")
	r.HandleFunc("/tradings/{tradingId}/items", createTradingItem(t)).
		Methods("POST")
	r.HandleFunc("/tradings/{tradingId}/items/{itemId}", updateTradingItem(t)).
		Methods("PUT")
	r.HandleFunc("/tradings/{tradingId}/items/{itemId}", deleteTradingItem(t)).
		Methods("DELETE")
	r.HandleFunc("/companies", getCompanies(c)).
		Methods("GET")
	r.HandleFunc("/companies", createCompany(c)).
		Methods("POST")
	r.HandleFunc("/companies/{companyId}", updateCompany(c)).
		Methods("PUT")
	r.HandleFunc("/sequences/{seqType}", getNextNumber(services)).
		Methods("POST")
	// Environment
	r.HandleFunc("/environments", getEnvironment(services)).
		Methods("GET")
	r.HandleFunc("/environments", saveEnvironment(services)).
		Methods("PUT")
	r.HandleFunc("/myCompany/name", getMyCompanyName(services)).
		Methods("GET")
}
