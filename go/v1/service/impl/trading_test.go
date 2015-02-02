package impl

import (
	"../../model"
	m "../../model/mock"
	"testing"
)

func TestTrading0000_GetListByUser(t *testing.T) {
	sessionDAO := &m.SessionDAO{
		GetByTokenResult: &model.Session{
			Token: "testToken",
		},
	}
	var list []*model.Trading
	list = append(list, &model.Trading{
		Id:         "trade1111",
		CompanyId:  "company2233",
		Subject:    "subject3344",
		WorkFrom:   1122,
		WorkTo:     3344,
		AssigneeId: "user2233",
		Product:    "product",
	})
	list = append(list, &model.Trading{
		Id: "trade2222",
	})
	tradingDAO := &m.TradingDAO{
		GetListByUserResult: list,
	}

	s := NewTradingSerivce(sessionDAO, tradingDAO)

	token := "token1122"
	r := s.GetListByUser(token)
	if r == nil {
		t.Errorf("Result must not be nil")
		return
	}
	if r.Status() != 200 {
		t.Errorf("Wrong status : %d", r.Status())
		return
	}

	json := json(r)
	tradings, _ := json.Array("tradings")
	if len(tradings) != 2 {
		t.Errorf("Wrong length : %d", len(tradings))
		return
	}
	item, _ := tradings.Object(0)
	if v, _ := item.String("id"); v != "trade1111" {
		t.Errorf("Wrong id : %s", v)
	}
	if v, _ := item.String("company_id"); v != "company2233" {
		t.Errorf("Wrong company id : %s", v)
	}
	if v, _ := item.String("subject"); v != "subject3344" {
		t.Errorf("Wrong subject : %s", v)
	}
	if v, _ := item.Int("work_from"); v != 1122 {
		t.Errorf("Wrong work from : %d", v)
	}
	if v, _ := item.Int("work_to"); v != 3344 {
		t.Errorf("Wrong work to : %d", v)
	}
	if v, _ := item.String("assignee"); v != "user2233" {
		t.Errorf("Wrong assignee : %d", v)
	}
}

func TestTrading0100_Create(t *testing.T) {
	sessionDAO := &m.SessionDAO{
		GetByTokenResult: &model.Session{
			Token: "testToken",
		},
	}
	tradingDAO := &m.TradingDAO{
		CreateResult: &model.Trading{
			Id:         "trade1111",
			CompanyId:  "company2233",
			Subject:    "subject3344",
			WorkFrom:   1122,
			WorkTo:     3344,
			AssigneeId: "user2233",
			Product:    "product",
		},
	}

	s := NewTradingSerivce(sessionDAO, tradingDAO)

	// params
	token := "token1122"
	date := "20150203"
	companyId := "company1122"
	subject := "subject3344"
	product := "product4455"
	workFrom := int64(100)
	workTo := int64(200)

	r := s.Create(token, date, companyId, subject, product, workFrom, workTo)
	if r == nil {
		t.Errorf("Result must not be nil")
		return
	}
	if r.Status() != 201 {
		t.Errorf("Wrong status : %d", r.Status())
		return
	}

	json := json(r)
	if v, _ := json.String("id"); v != "trade1111" {
		t.Errorf("Wrong id : %s", v)
	}
}
