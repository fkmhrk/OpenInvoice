package rest

import (
	s "../service"
	rj "github.com/fkmhrk-go/rawjson"
	"github.com/gorilla/mux"
	"net/http"
)

func getTradings(trading s.TradingService) handler {
	return makeHandler(func(token, tType string,
		req *http.Request) s.Result {
		return trading.GetListByUser(token)
	})
}

func createTrading(trading s.TradingService) handler {
	return makeJsonHandler(func(token, tType string,
		json rj.RawJsonObject) s.Result {
		// read input
		date, _ := json.String("date")
		companyId, _ := json.String("company_id")
		titleType, _ := json.Int("title_type")
		subject, _ := json.String("subject")
		workFrom, _ := json.Long("work_from")
		workTo, _ := json.Long("work_to")
		quotationDate, _ := json.Long("quotation_date")
		billDate, _ := json.Long("bill_date")
		taxRate, _ := json.Float("tax_rate")
		product, _ := json.String("product")

		return trading.Create(token, date, companyId,
			subject, product, titleType, workFrom, workTo,
			quotationDate, billDate, float32(taxRate))
	})
}

func updateTrading(trading s.TradingService) handler {
	return makeHandler(func(token, tType string,
		req *http.Request) s.Result {
		// read path param
		vars := mux.Vars(req)
		tradingId := vars["tradingId"]

		// to json
		json, _ := rj.ObjectFromString(readBody(req))

		// read input
		companyId, _ := json.String("company_id")
		subject, _ := json.String("subject")
		titleType, _ := json.Int("title_type")
		workFrom, _ := json.Long("work_from")
		workTo, _ := json.Long("work_to")
		quotationDate, _ := json.Long("quotation_date")
		billDate, _ := json.Long("bill_date")
		taxRate, _ := json.Float("tax_rate")
		product, _ := json.String("product")

		return trading.Update(token, tradingId, companyId,
			subject, product, titleType, workFrom, workTo,
			quotationDate, billDate, float32(taxRate))
	})
}

func getTradingItems(trading s.TradingService) handler {
	return makeHandler(func(token, tType string,
		req *http.Request) s.Result {
		// read path param
		vars := mux.Vars(req)
		tradingId := vars["tradingId"]

		return trading.GetItemListByTradingId(token, tradingId)
	})
}

func createTradingItem(trading s.TradingService) handler {
	return makeHandler(func(token, tType string,
		req *http.Request) s.Result {
		// read path param
		vars := mux.Vars(req)
		tradingId := vars["tradingId"]

		// to json
		json, _ := rj.ObjectFromString(readBody(req))

		// get values
		sortOrder, _ := json.Int("sort_order")
		subject, _ := json.String("subject")
		unitPrice, _ := json.Int("unit_price")
		amount, _ := json.Int("amount")
		degree, _ := json.String("degree")
		taxType, _ := json.Int("tax_type")
		memo, _ := json.String("memo")

		return trading.CreateItem(token, tradingId, subject, degree, memo, sortOrder, unitPrice, amount, taxType)
	})
}

func updateTradingItem(trading s.TradingService) handler {
	return makeHandler(func(token, tType string,
		req *http.Request) s.Result {
		// read path param
		vars := mux.Vars(req)
		tradingId := vars["tradingId"]
		id := vars["itemId"]

		// to json
		json, _ := rj.ObjectFromString(readBody(req))

		// get values
		sortOrder, _ := json.Int("sort_order")
		subject, _ := json.String("subject")
		unitPrice, _ := json.Int("unit_price")
		amount, _ := json.Int("amount")
		degree, _ := json.String("degree")
		taxType, _ := json.Int("tax_type")
		memo, _ := json.String("memo")

		return trading.UpdateItem(token, id, tradingId, subject, degree, memo, sortOrder, unitPrice, amount, taxType)
	})
}
