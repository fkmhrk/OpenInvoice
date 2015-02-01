package rest

import (
	s "../service"
	"fmt"
	rj "github.com/fkmhrk-go/rawjson"
	"net/http"
)

//func getToken() func(http.ResponseWriter, *http.Request) {
func getToken(user s.UserService) handler {
	return func(w http.ResponseWriter, req *http.Request) {
		// read input
		json, _ := rj.ObjectFromString(readBody(req))
		// get values
		username, _ := json.String("username")
		password, _ := json.String("password")
		// execute
		result := user.GetToken(username, password)
		for k, v := range result.Headers() {
			w.Header().Set(k, v)
		}

		w.WriteHeader(result.Status())
		fmt.Fprintf(w, result.Body())
	}
}
