package rest

import (
	s "../service"
	"fmt"
	rj "github.com/fkmhrk-go/rawjson"
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

func createTrading(trading s.TradingService) handler {
	return func(w http.ResponseWriter, req *http.Request) {
		// read header
		authorization := req.Header.Get("authorization")
		_, token := parseAuth(authorization)

		// read input
		json, _ := rj.ObjectFromString(readBody(req))
		date, _ := json.String("date")
		companyId, _ := json.String("company_id")
		subject, _ := json.String("subject")
		workFrom, _ := json.Long("work_from")
		workTo, _ := json.Long("work_to")
		product, _ := json.String("product")

		result := trading.Create(token, date, companyId, subject, product, workFrom, workTo)
		for k, v := range result.Headers() {
			w.Header().Set(k, v)
		}

		w.WriteHeader(result.Status())
		fmt.Fprintf(w, result.Body())
	}
}
