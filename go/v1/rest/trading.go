package rest

import (
	m "../model"
	s "../service"
	rj "github.com/fkmhrk-go/rawjson"
	"github.com/gorilla/mux"
	"net/http"
)

func getTradings(trading s.TradingService) http.HandlerFunc {
	return makeBaseHandler(func(req *http.Request) s.Result {
		return trading.GetListByUser()
	})
}

func createTrading(trading s.TradingService) handler {
	return makeJsonHandler(func(token, tType string,
		json rj.RawJsonObject) s.Result {
		// read input
		companyId, _ := json.String("company_id")
		titleType, _ := json.Int("title_type")
		subject, _ := json.String("subject")
		workFrom, _ := json.Long("work_from")
		workTo, _ := json.Long("work_to")
		total, _ := json.Long("total")
		quotationDate, _ := json.Long("quotation_date")
		billDate, _ := json.Long("bill_date")
		deliveryDate, _ := json.Long("delivery_date")
		taxRate, _ := json.Float("tax_rate")
		product, _ := json.String("product")
		memo, _ := json.String("memo")

		return trading.Create(token, companyId,
			subject, product, memo, titleType, workFrom, workTo,
			total, quotationDate, billDate, deliveryDate, float32(taxRate))
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
		total, _ := json.Long("total")
		quotationDate, _ := json.Long("quotation_date")
		quotationNumber, _ := json.String("quotation_number")
		billDate, _ := json.Long("bill_date")
		billNumber, _ := json.String("bill_number")
		deliveryDate, _ := json.Long("delivery_date")
		deliveryNumber, _ := json.String("delivery_number")
		taxRate, _ := json.Float("tax_rate")
		assignee, _ := json.String("assignee")
		product, _ := json.String("product")
		memo, _ := json.String("memo")

		return trading.Update(token, s.Trading{
			m.Trading{
				Id:              tradingId,
				CompanyId:       companyId,
				Subject:         subject,
				TitleType:       titleType,
				WorkFrom:        workFrom,
				WorkTo:          workTo,
				QuotationDate:   quotationDate,
				QuotationNumber: quotationNumber,
				BillDate:        billDate,
				BillNumber:      billNumber,
				DeliveryDate:    deliveryDate,
				DeliveryNumber:  deliveryNumber,
				TaxRate:         float32(taxRate),
				AssigneeId:      assignee,
				Product:         product,
				Memo:            memo,
				Total:           total,
			},
		})
	})
}

func deleteTrading(services s.Services) handler {
	return makeHandler(func(token, tType string, req *http.Request) s.Result {
		// read path param
		vars := mux.Vars(req)
		tradingId := vars["tradingId"]

		return services.Trading.Delete(token, tradingId)
	})
}

func getTradingItems(trading s.TradingService) http.HandlerFunc {
	return makeBaseHandler(func(req *http.Request) s.Result {
		// read path param
		vars := mux.Vars(req)
		tradingId := vars["tradingId"]

		return trading.GetItemListByTradingId(tradingId)
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
		amount, _ := json.Float("amount")
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
		amount, _ := json.Float("amount")
		degree, _ := json.String("degree")
		taxType, _ := json.Int("tax_type")
		memo, _ := json.String("memo")

		return trading.UpdateItem(token, id, tradingId, subject, degree, memo, sortOrder, unitPrice, amount, taxType)
	})
}

func deleteTradingItem(trading s.TradingService) handler {
	return makeHandler(func(token, tType string,
		req *http.Request) s.Result {
		// read path param
		vars := mux.Vars(req)
		tradingId := vars["tradingId"]
		id := vars["itemId"]

		return trading.DeleteItem(token, id, tradingId)
	})
}
