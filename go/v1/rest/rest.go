package rest

import (
	s "../service"
	"github.com/gorilla/mux"
	"net/http"
)

type handler func(http.ResponseWriter, *http.Request)

func SetHandlers(r *mux.Router, u s.UserService, t s.TradingService) {
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
}
