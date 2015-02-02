package rest

import (
	s "../service"
	"fmt"
	//rj "github.com/fkmhrk-go/rawjson"
	"net/http"
)

func getTradings(trading s.TradingService) handler {
	return func(w http.ResponseWriter, req *http.Request) {
		// read header
		authorization := req.Header.Get("authorization")
		_, token := parseAuth(authorization)

		result := trading.GetListByUser(token)
		for k, v := range result.Headers() {
			w.Header().Set(k, v)
		}

		w.WriteHeader(result.Status())
		fmt.Fprintf(w, result.Body())
	}
}
