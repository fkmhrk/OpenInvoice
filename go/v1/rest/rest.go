package rest

import (
	s "../service"
	"github.com/gorilla/mux"
	"net/http"
)

type handler func(http.ResponseWriter, *http.Request)

func SetHandlers(r *mux.Router, u s.UserService) {
	r.HandleFunc("/token", getToken(u)).
		Methods("POST")
}
