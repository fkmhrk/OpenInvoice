package impl

import (
	"net/http"
	"strconv"
	"time"

	m "github.com/fkmhrk/OpenInvoice/v1/model"
	"github.com/fkmhrk/OpenInvoice/v1/model/response"
	ss "github.com/fkmhrk/OpenInvoice/v1/service/trading"
	"github.com/mokelab-go/server/entity"
)

type tradingService struct {
	sessionDAO m.SessionDAO
	tradingDAO m.TradingDAO
	envDAO     m.EnvDAO
	seqDAO     m.SeqDAO
}

func New(s m.SessionDAO, t m.TradingDAO, models *m.Models) *tradingService {
	return &tradingService{
		sessionDAO: s,
		tradingDAO: t,
		envDAO:     models.Env,
		seqDAO:     models.Seq,
	}
}

func (s *tradingService) GetListByUser() entity.Response {
	// get : fixed we use GetList
	tradings, err := s.tradingDAO.GetList()
	if err != nil {
		return response.Error(http.StatusInternalServerError, response.MSG_SERVER_ERROR)
	}
	list := make([]interface{}, 0)
	for _, t := range tradings {
		list = append(list, s.toJson(t))
	}
	body := map[string]interface{}{
		"tradings": list,
	}
	return entity.Response{
		Status: http.StatusOK,
		Body:   body,
	}
}

func (s *tradingService) GetTradingByID(id string) entity.Response {
	// get : fixed we use GetList
	trading, err := s.tradingDAO.GetById(id)
	if err != nil {
		return response.Error(http.StatusInternalServerError, response.MSG_SERVER_ERROR)
	}
	body := s.toJson(trading)
	return entity.Response{
		Status: http.StatusOK,
		Body:   body,
	}
}

func (s *tradingService) Create(session *m.Session, companyId, subject, product, memo string, titleType int, workFrom, workTo, total, quotationDate, billDate, deliveryDate int64, taxRate float32) entity.Response {
	// input check
	if len(companyId) == 0 {
		return response.Error(http.StatusBadRequest, response.MSG_ERR_COMPANY_ID_EMPTY)
	}
	if len(subject) == 0 {
		return response.Error(http.StatusBadRequest, response.MSG_ERR_SUBJECT_EMPTY)
	}

	// create
	item, err := s.tradingDAO.Create(companyId, subject, titleType, workFrom, workTo, total, quotationDate, billDate, deliveryDate, taxRate, session.UserId, product, memo)
	if err != nil {
		return response.Error(http.StatusInternalServerError, response.MSG_SERVER_ERROR)
	}

	body := map[string]interface{}{
		"id": item.Id,
	}
	return entity.Response{
		Status: http.StatusCreated,
		Body:   body,
	}
}

func (s *tradingService) Update(trading ss.Trading) entity.Response {
	// input check
	if len(trading.Id) == 0 {
		return response.Error(http.StatusBadRequest, response.MSG_ERR_ID_EMPTY)
	}
	if len(trading.CompanyId) == 0 {
		return response.Error(http.StatusBadRequest, response.MSG_ERR_COMPANY_ID_EMPTY)
	}
	if len(trading.Subject) == 0 {
		return response.Error(http.StatusBadRequest, response.MSG_ERR_SUBJECT_EMPTY)
	}

	// get item
	item, err := s.tradingDAO.GetById(trading.Id)
	if err != nil {
		return response.Error(http.StatusInternalServerError, response.MSG_SERVER_ERROR)
	}
	if item == nil {
		return response.Error(http.StatusNotFound, response.MSG_TRADING_NOT_FOUND)
	}
	// update
	item2, err := s.tradingDAO.Update(trading.Trading)
	if err != nil {
		return response.Error(http.StatusInternalServerError, response.MSG_SERVER_ERROR)
	}

	body := map[string]interface{}{
		"id": item2.Id,
	}
	return entity.Response{
		Status: http.StatusOK,
		Body:   body,
	}
}

func (o *tradingService) Delete(tradingId string) entity.Response {
	// input check
	if len(tradingId) == 0 {
		return response.Error(http.StatusBadRequest, response.MSG_ERR_ID_EMPTY)
	}

	// delete
	err := o.tradingDAO.Delete(tradingId)
	if err != nil {
		return response.Error(http.StatusInternalServerError, response.MSG_SERVER_ERROR)
	}

	return entity.Response{
		Status: http.StatusNoContent,
	}
}

func (s *tradingService) GetItemListByTradingId(tradingId string) entity.Response {
	// get trading items
	items, err := s.tradingDAO.GetItemsById(tradingId)
	if err != nil {
		return response.Error(http.StatusInternalServerError, response.MSG_SERVER_ERROR)
	}
	list := make([]interface{}, 0)
	for _, t := range items {
		list = append(list, map[string]interface{}{
			"id":         t.Id,
			"subject":    t.Subject,
			"unit_price": t.UnitPrice,
			"amount":     t.Amount,
			"degree":     t.Degree,
			"tax_type":   t.TaxType,
			"memo":       t.Memo,
		})
	}
	body := map[string]interface{}{
		"items": list,
	}
	return entity.Response{
		Status: http.StatusOK,
		Body:   body,
	}
}

func (s *tradingService) CreateItem(tradingId, subject, degree, memo string, sortOrder, unitPrice int, amount float64, taxType int) entity.Response {
	// create
	item, err := s.tradingDAO.CreateItem(tradingId, subject, degree, memo, sortOrder, unitPrice, amount, taxType)
	if err != nil {
		return response.Error(http.StatusInternalServerError, response.MSG_SERVER_ERROR)
	}
	body := map[string]interface{}{
		"id": item.Id,
	}
	return entity.Response{
		Status: http.StatusCreated,
		Body:   body,
	}
}

func (s *tradingService) UpdateItem(id, tradingId, subject, degree, memo string, sortOrder, unitPrice int, amount float64, taxType int) entity.Response {
	// input check
	// Update
	item, err := s.tradingDAO.UpdateItem(id, tradingId, subject, degree, memo, sortOrder, unitPrice, amount, taxType)
	if err != nil {
		return response.Error(http.StatusInternalServerError, response.MSG_SERVER_ERROR)
	}
	body := map[string]interface{}{
		"id": item.Id,
	}
	return entity.Response{
		Status: http.StatusOK,
		Body:   body,
	}
}

func (s *tradingService) DeleteItem(id, tradingId string) entity.Response {
	// soft delete
	err := s.tradingDAO.SoftDeleteItem(id, tradingId)
	if err != nil {
		return response.Error(http.StatusInternalServerError, response.MSG_SERVER_ERROR)
	}
	return entity.Response{
		Status: http.StatusNoContent,
	}
}

func (o *tradingService) GetNextNumber(seqType string, date int64) entity.Response {
	var seqTypeInt m.SeqType
	switch seqType {
	case "quotation":
		seqTypeInt = m.SeqType_Quotation
		break
	case "delivery":
		seqTypeInt = m.SeqType_Delivery
		break
	case "bill":
		seqTypeInt = m.SeqType_Bill
		break
	default:
		return response.Error(http.StatusBadRequest, response.MSG_INVALID_SEQUENCE_TYPE)
	}
	// determine year
	t := time.Unix(date/1000, 0)
	year := t.Year()
	month := t.Month()
	env, err := o.envDAO.Get("closing_month")
	if err != nil {
		return response.Error(http.StatusInternalServerError, response.MSG_SERVER_ERROR)
	}
	intVal, _ := strconv.Atoi(env.Value)
	if int(month) <= intVal {
		year--
	}
	// get next sequence
	seq, err := o.seqDAO.Next(seqTypeInt, year)
	if err != nil {
		return response.Error(http.StatusInternalServerError, response.MSG_SERVER_ERROR)
	}

	body := map[string]interface{}{
		"number": year*10000 + seq.Value,
	}
	return entity.Response{
		Status: http.StatusOK,
		Body:   body,
	}
}

func (s *tradingService) toJson(t *m.Trading) map[string]interface{} {
	return map[string]interface{}{
		"id":               t.Id,
		"company_id":       t.CompanyId,
		"subject":          t.Subject,
		"title_type":       t.TitleType,
		"work_from":        t.WorkFrom,
		"work_to":          t.WorkTo,
		"total":            t.Total,
		"quotation_date":   t.QuotationDate,
		"quotation_number": t.QuotationNumber,
		"bill_date":        t.BillDate,
		"bill_number":      t.BillNumber,
		"delivery_date":    t.DeliveryDate,
		"delivery_number":  t.DeliveryNumber,
		"tax_rate":         t.TaxRate,
		"assignee":         t.AssigneeId,
		"product":          t.Product,
		"memo":             t.Memo,
		"modified_time":    t.ModifiedTime * 1000,
	}
}
