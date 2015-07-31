package impl

import (
	s "../"
	m "../../model"
	"strconv"
	"time"
)

type tradingService struct {
	sessionDAO m.SessionDAO
	tradingDAO m.TradingDAO
	envDAO     m.EnvDAO
	seqDAO     m.SeqDAO
}

func NewTradingSerivce(s m.SessionDAO, t m.TradingDAO, models *m.Models) *tradingService {
	return &tradingService{
		sessionDAO: s,
		tradingDAO: t,
		envDAO:     models.Env,
		seqDAO:     models.Seq,
	}
}

func (s *tradingService) GetListByUser(token string) s.Result {
	// input check
	session, err := s.sessionDAO.GetByToken(token)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	if session == nil {
		return errorResult(401, MSG_WRONG_TOKEN)
	}
	// get : fixed we use GetList
	tradings, err := s.tradingDAO.GetList()
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	list := make([]interface{}, 0)
	for _, t := range tradings {
		list = append(list, s.toJson(t))
	}
	body := map[string]interface{}{
		"tradings": list,
	}
	return jsonResult(200, body)
}

func (s *tradingService) Create(token, companyId, subject, product, memo string, titleType int, workFrom, workTo, total, quotationDate, billDate, deliveryDate int64, taxRate float32) s.Result {
	// input check
	if len(companyId) == 0 {
		return errorResult(400, MSG_ERR_COMPANY_ID_EMPTY)
	}
	if len(subject) == 0 {
		return errorResult(400, MSG_ERR_SUBJECT_EMPTY)
	}

	// get session
	session, err := s.sessionDAO.GetByToken(token)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	if session == nil {
		return errorResult(401, MSG_WRONG_TOKEN)
	}
	// create
	item, err := s.tradingDAO.Create(companyId, subject, titleType, workFrom, workTo, total, quotationDate, billDate, deliveryDate, taxRate, session.UserId, product, memo)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}

	body := map[string]interface{}{
		"id": item.Id,
	}
	return jsonResult(201, body)
}

func (s *tradingService) Update(token string, trading s.Trading) s.Result {
	// input check
	if len(trading.Id) == 0 {
		return errorResult(400, MSG_ERR_ID_EMPTY)
	}
	if len(trading.CompanyId) == 0 {
		return errorResult(400, MSG_ERR_COMPANY_ID_EMPTY)
	}
	if len(trading.Subject) == 0 {
		return errorResult(400, MSG_ERR_SUBJECT_EMPTY)
	}

	// get session
	session, err := s.sessionDAO.GetByToken(token)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	if session == nil {
		return errorResult(401, MSG_WRONG_TOKEN)
	}

	// get item
	item, err := s.tradingDAO.GetById(trading.Id)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	if item == nil {
		return errorResult(404, MSG_TRADING_NOT_FOUND)
	}
	// update
	item2, err := s.tradingDAO.Update(trading.Trading)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}

	body := map[string]interface{}{
		"id": item2.Id,
	}
	return jsonResult(200, body)
}

func (o *tradingService) Delete(token, tradingId string) s.Result {
	// input check
	if len(tradingId) == 0 {
		return errorResult(400, MSG_ERR_ID_EMPTY)
	}

	// get session
	session, err := o.sessionDAO.GetByToken(token)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	if session == nil {
		return errorResult(401, MSG_WRONG_TOKEN)
	}

	// delete
	err = o.tradingDAO.Delete(tradingId)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}

	return &result{
		status: 204,
		body:   "",
	}
}

func (s *tradingService) GetItemListByTradingId(token, tradingId string) s.Result {
	// input check
	session, err := s.sessionDAO.GetByToken(token)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	if session == nil {
		return errorResult(401, MSG_WRONG_TOKEN)
	}
	// get trading
	// get
	items, err := s.tradingDAO.GetItemsById(tradingId)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
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
	return jsonResult(200, body)
}

func (s *tradingService) CreateItem(token, tradingId, subject, degree, memo string, sortOrder, unitPrice, amount, taxType int) s.Result {
	// input check
	// get session
	session, err := s.sessionDAO.GetByToken(token)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	if session == nil {
		return errorResult(401, MSG_WRONG_TOKEN)
	}
	// get trading
	// create
	item, err := s.tradingDAO.CreateItem(tradingId, subject, degree, memo, sortOrder, unitPrice, amount, taxType)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	body := map[string]interface{}{
		"id": item.Id,
	}
	return jsonResult(201, body)
}

func (s *tradingService) UpdateItem(token, id, tradingId, subject, degree, memo string, sortOrder, unitPrice, amount, taxType int) s.Result {
	// input check
	// get session
	session, err := s.sessionDAO.GetByToken(token)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	if session == nil {
		return errorResult(401, MSG_WRONG_TOKEN)
	}
	// get trading
	// create
	item, err := s.tradingDAO.UpdateItem(id, tradingId, subject, degree, memo, sortOrder, unitPrice, amount, taxType)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	body := map[string]interface{}{
		"id": item.Id,
	}
	return jsonResult(200, body)
}

func (s *tradingService) DeleteItem(token, id, tradingId string) s.Result {
	// input check
	// get session
	session, err := s.sessionDAO.GetByToken(token)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	if session == nil {
		return errorResult(401, MSG_WRONG_TOKEN)
	}
	// soft delete
	err = s.tradingDAO.SoftDeleteItem(id, tradingId)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	return &result{
		status: 204,
		body:   "",
	}
}

func (o *tradingService) GetNextNumber(token, seqType string, date int64) s.Result {
	// input check
	if isEmpty(token) {
		return errorResult(401, s.ERR_TOKEN_EMPTY)
	}
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
		return errorResult(400, s.ERR_INVALID_SEQUENCE_TYPE)
	}
	// get session
	session, err := o.sessionDAO.GetByToken(token)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	if session == nil {
		return errorResult(401, MSG_WRONG_TOKEN)
	}
	// determine year
	t := time.Unix(date/1000, 0)
	year := t.Year()
	month := t.Month()
	env, err := o.envDAO.Get("closing_month")
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	intVal, _ := strconv.Atoi(env.Value)
	if int(month) <= intVal {
		year--
	}
	// get next sequence
	seq, err := o.seqDAO.Next(seqTypeInt, year)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}

	body := map[string]interface{}{
		"number": year*10000 + seq.Value,
	}
	return jsonResult(200, body)
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
