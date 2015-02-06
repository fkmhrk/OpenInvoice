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
	r.HandleFunc("/tradings", getTradings(t)).
		Methods("GET")
	r.HandleFunc("/tradings", createTrading(t)).
		Methods("POST")
	r.HandleFunc("/tradings/{tradingId}", getTradingItems(t)).
		Methods("GET")
	r.HandleFunc("/tradings/{tradingId}", createTradingItem(t)).
		Methods("POST")
}
