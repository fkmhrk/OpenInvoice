package rest

import (
	"net/http"

	rj "github.com/fkmhrk-go/rawjson"
	"github.com/fkmhrk/OpenInvoice/v1/entity"
	"github.com/fkmhrk/OpenInvoice/v1/rest/service"
	"github.com/mokelab-go/hop"
)

func getTradings(trading service.Trading) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		resp := trading.GetListByUser()
		resp.Write(w)
	}
}

func getTrading(trading service.Trading) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		c := req.Context()
		tradingId := hop.PathString(c, "tradingId")
		resp := trading.GetTradingByID(tradingId)
		resp.Write(w)
	}
}

func createTrading(trading service.Trading) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// read input
		c := req.Context()
		session := session(c)
		json := rj.RawJsonObject(hop.BodyJSON(c))

		companyID, _ := json.String("company_id")
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

		resp := trading.Create(session, companyID,
			subject, product, memo, titleType, workFrom, workTo,
			total, quotationDate, billDate, deliveryDate, float32(taxRate))
		resp.Write(w)

	}
}

func updateTrading(tradingService service.Trading) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		c := req.Context()
		tradingID := hop.PathString(c, "tradingId")

		// to json
		json := rj.RawJsonObject(hop.BodyJSON(c))

		// read input
		companyID, _ := json.String("company_id")
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

		resp := tradingService.Update(service.TradingData{
			entity.Trading{
				Id:              tradingID,
				CompanyId:       companyID,
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
		resp.Write(w)
	}

}

func deleteTrading(services service.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// read path param
		tradingID := hop.PathString(req.Context(), "tradingId")

		resp := services.Trading.Delete(tradingID)
		resp.Write(w)
	}
}

func getTradingItems(trading service.Trading) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// read path param
		tradingID := hop.PathString(req.Context(), "tradingId")

		resp := trading.GetItemListByTradingId(tradingID)
		resp.Write(w)
	}
}

func createTradingItem(trading service.Trading) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		c := req.Context()
		tradingID := hop.PathString(c, "tradingId")

		// to json
		json := rj.RawJsonObject(hop.BodyJSON(c))

		// get values
		sortOrder, _ := json.Int("sort_order")
		subject, _ := json.String("subject")
		unitPrice, _ := json.Int("unit_price")
		amount, _ := json.Float("amount")
		degree, _ := json.String("degree")
		taxType, _ := json.Int("tax_type")
		memo, _ := json.String("memo")

		resp := trading.CreateItem(tradingID, subject, degree, memo, sortOrder, unitPrice, amount, taxType)
		resp.Write(w)
	}
}

func updateTradingItem(trading service.Trading) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// read path param
		c := req.Context()
		tradingID := hop.PathString(c, "tradingId")
		id := hop.PathString(c, "itemId")

		// to json
		json := rj.RawJsonObject(hop.BodyJSON(c))

		// get values
		sortOrder, _ := json.Int("sort_order")
		subject, _ := json.String("subject")
		unitPrice, _ := json.Int("unit_price")
		amount, _ := json.Float("amount")
		degree, _ := json.String("degree")
		taxType, _ := json.Int("tax_type")
		memo, _ := json.String("memo")

		resp := trading.UpdateItem(id, tradingID, subject, degree, memo, sortOrder, unitPrice, amount, taxType)
		resp.Write(w)
	}
}

func deleteTradingItem(trading service.Trading) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// read path param
		c := req.Context()
		tradingID := hop.PathString(c, "tradingId")
		id := hop.PathString(c, "itemId")

		resp := trading.DeleteItem(id, tradingID)
		resp.Write(w)
	}
}
