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
