package impl

import (
	s "../"
	m "../../model"
)

type tradingService struct {
	sessionDAO m.SessionDAO
	tradingDAO m.TradingDAO
}

func NewTradingSerivce(s m.SessionDAO, t m.TradingDAO) *tradingService {
	return &tradingService{
		sessionDAO: s,
		tradingDAO: t,
	}
}

func (s *tradingService) GetListByUser(token string) s.Result {
	// input check
	session, err := s.sessionDAO.GetByToken(token)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	if session == nil {
		return errorResult(400, MSG_WRONG_TOKEN)
	}
	// get
	tradings, err := s.tradingDAO.GetListByUser(session.UserId)
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

func (s *tradingService) Create(token, date, companyId, subject, product string, titleType int, workFrom, workTo, total, quotationDate, billDate int64, taxRate float32) s.Result {
	// input check
	if len(date) == 0 {
		return errorResult(400, MSG_ERR_DATE_EMPTY)
	}
	if len(companyId) == 0 {
		return errorResult(400, MSG_ERR_COMPANY_ID_EMPTY)
	}
	if len(subject) == 0 {
		return errorResult(400, MSG_ERR_SUBJECT_EMPTY)
	}
	if len(product) == 0 {
		return errorResult(400, MSG_ERR_PRODUCT_EMPTY)
	}

	// get session
	session, err := s.sessionDAO.GetByToken(token)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	if session == nil {
		return errorResult(400, MSG_WRONG_TOKEN)
	}
	// create
	item, err := s.tradingDAO.Create(date, companyId, subject, titleType, workFrom, workTo, total, quotationDate, billDate, taxRate, session.UserId, product)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}

	body := map[string]interface{}{
		"id": item.Id,
	}
	return jsonResult(201, body)
}

func (s *tradingService) Update(token, id, companyId, subject, product string, titleType int, workFrom, workTo, quotationDate, billDate int64, taxRate float32) s.Result {
	// input check
	if len(id) == 0 {
		return errorResult(400, MSG_ERR_ID_EMPTY)
	}
	if len(companyId) == 0 {
		return errorResult(400, MSG_ERR_COMPANY_ID_EMPTY)
	}
	if len(subject) == 0 {
		return errorResult(400, MSG_ERR_SUBJECT_EMPTY)
	}
	if len(product) == 0 {
		return errorResult(400, MSG_ERR_PRODUCT_EMPTY)
	}

	// get session
	session, err := s.sessionDAO.GetByToken(token)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	if session == nil {
		return errorResult(400, MSG_WRONG_TOKEN)
	}

	// get item
	item, err := s.tradingDAO.GetById(id, session.UserId)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	if item == nil {
		return errorResult(404, MSG_TRADING_NOT_FOUND)
	}
	// update
	item2, err := s.tradingDAO.Update(id, companyId, subject, titleType, workFrom, workTo, quotationDate, billDate, taxRate, session.UserId, product)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}

	body := map[string]interface{}{
		"id": item2.Id,
	}
	return jsonResult(200, body)
}

func (s *tradingService) GetItemListByTradingId(token, tradingId string) s.Result {
	// input check
	session, err := s.sessionDAO.GetByToken(token)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	if session == nil {
		return errorResult(400, MSG_WRONG_TOKEN)
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
		return errorResult(400, MSG_WRONG_TOKEN)
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
		return errorResult(400, MSG_WRONG_TOKEN)
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
		return errorResult(400, MSG_WRONG_TOKEN)
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

func (s *tradingService) toJson(t *m.Trading) map[string]interface{} {
	return map[string]interface{}{
		"id":               t.Id,
		"company_id":       t.CompanyId,
		"title_type":       t.TitleType,
		"subject":          t.Subject,
		"work_from":        t.WorkFrom,
		"work_to":          t.WorkTo,
		"total":            t.Total,
		"quotation_date":   t.QuotationDate,
		"quotation_number": t.QuotationNumber,
		"bill_date":        t.BillDate,
		"bill_number":      t.BillNumber,
		"tax_rate":         t.TaxRate,
		"assignee":         t.AssigneeId,
		"product":          t.Product,
	}
}
