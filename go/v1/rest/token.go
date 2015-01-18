package rest

import (
	"fmt"
	"net/http"
)

//func getToken() func(http.ResponseWriter, *http.Request) {
func getToken() handler {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "{\"id\":\"fkm\",\"access_token\":\"token1234\"}")
	}
}
