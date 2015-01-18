package rest

import (
	"github.com/gorilla/mux"
	"net/http"
)

type handler func(http.ResponseWriter, *http.Request)

func SetHandlers(r *mux.Router) {
	r.HandleFunc("/token", getToken()).
		Methods("POST")
}
