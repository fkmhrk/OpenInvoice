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
		list = append(list, map[string]interface{}{
			"id":         t.Id,
			"company_id": t.CompanyId,
			"subject":    t.Subject,
			"work_from":  t.WorkFrom,
			"work_to":    t.WorkTo,
			"assignee":   t.AssigneeId,
			"product":    t.Product,
		})
	}
	body := map[string]interface{}{
		"tradings": list,
	}
	return jsonResult(200, body)
}

func (s *tradingService) Create(token, date, companyId, subject, product string, workFrom, workTo int64) s.Result {
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
	item, err := s.tradingDAO.Create(date, companyId, subject, workFrom, workTo, session.UserId, product)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}

	body := map[string]interface{}{
		"id": item.Id,
	}
	return jsonResult(201, body)
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
