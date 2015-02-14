package impl

import (
	"../../model"
	m "../../model/mock"
	"fmt"
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
		TitleType:  1,
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
	if v, _ := item.Int("title_type"); v != 1 {
		t.Errorf("Wrong title_type : %d", v)
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
	titleType := 1
	subject := "subject3344"
	product := "product4455"
	workFrom := int64(100)
	workTo := int64(200)
	quotationDate := int64(300)
	billDate := int64(400)
	taxRate := float32(8)

	r := s.Create(token, date, companyId, subject, product, titleType, workFrom, workTo, quotationDate, billDate, taxRate)
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

func TestTrading0200_Update(t *testing.T) {
	sessionDAO := &m.SessionDAO{
		GetByTokenResult: &model.Session{
			Token:  "testToken",
			UserId: "user1122",
		},
	}
	tradingDAO := &m.TradingDAO{
		GetByIdResult: &model.Trading{
			Id:         "trade1111",
			CompanyId:  "company2233",
			Subject:    "subject3344",
			WorkFrom:   2233,
			WorkTo:     4455,
			AssigneeId: "user2233",
			Product:    "product",
		},
		UpdateResult: &model.Trading{
			Id:         "trade1111",
			CompanyId:  "company2233",
			Subject:    "subject3344",
			WorkFrom:   2233,
			WorkTo:     4455,
			AssigneeId: "user2233",
			Product:    "product",
		},
	}

	s := NewTradingSerivce(sessionDAO, tradingDAO)

	// params
	token := "token1122"
	id := "20150203"
	companyId := "company1122"
	subject := "subject3344"
	product := "product4455"
	titleType := 1
	workFrom := int64(100)
	workTo := int64(200)
	quotationDate := int64(300)
	billDate := int64(400)
	taxRate := float32(8)

	r := s.Update(token, id, companyId, subject, product, titleType, workFrom, workTo, quotationDate, billDate, taxRate)
	if r == nil {
		t.Errorf("Result must not be nil")
		return
	}
	if r.Status() != 200 {
		t.Errorf("Wrong status : %d", r.Status())
		return
	}

	json := json(r)
	if v, _ := json.String("id"); v != "trade1111" {
		t.Errorf("Wrong id : %s", v)
	}
}

func TestTrading0200_GetItemsByTradingId(t *testing.T) {
	sessionDAO := &m.SessionDAO{
		GetByTokenResult: &model.Session{
			Token: "testToken",
		},
	}
	var list []*model.TradingItem
	for i := 0; i < 2; i++ {
		list = append(list, &model.TradingItem{
			Id:        fmt.Sprintf("trade%d", i),
			Subject:   fmt.Sprintf("subject%d", i),
			UnitPrice: i*100 + 100,
			Amount:    i*3 + 3,
			Degree:    fmt.Sprintf("D%d", i),
			TaxType:   1,
			Memo:      fmt.Sprintf("memo%d", i),
		})
	}
	tradingDAO := &m.TradingDAO{
		GetItemsByIdResult: list,
	}

	s := NewTradingSerivce(sessionDAO, tradingDAO)

	token := "token1122"
	tradingId := "tradingId1"
	r := s.GetItemListByTradingId(token, tradingId)
	if r == nil {
		t.Errorf("Result must not be nil")
		return
	}
	if r.Status() != 200 {
		t.Errorf("Wrong status : %d", r.Status())
		return
	}

	json := json(r)
	items, _ := json.Array("items")
	if len(items) != 2 {
		t.Errorf("Wrong length : %d", len(items))
		return
	}
	item, _ := items.Object(0)
	if v, _ := item.String("id"); v != "trade0" {
		t.Errorf("Wrong id : %s", v)
	}
	if v, _ := item.String("subject"); v != "subject0" {
		t.Errorf("Wrong subject : %s", v)
	}
	if v, _ := item.Int("unit_price"); v != 100 {
		t.Errorf("Wrong unit_price : %d", v)
	}
	if v, _ := item.Int("amount"); v != 3 {
		t.Errorf("Wrong amount : %d", v)
	}
	if v, _ := item.String("degree"); v != "D0" {
		t.Errorf("Wrong degree : %s", v)
	}
	if v, _ := item.Int("tax_type"); v != 1 {
		t.Errorf("Wrong tax_type : %d", v)
	}
	if v, _ := item.String("memo"); v != "memo0" {
		t.Errorf("Wrong memo : %s", v)
	}
}

func TestTrading0300_CreateItem(t *testing.T) {
	sessionDAO := &m.SessionDAO{
		GetByTokenResult: &model.Session{
			Token: "testToken",
		},
	}
	tradingDAO := &m.TradingDAO{
		CreateItemResult: &model.TradingItem{
			Id: "item2233",
		},
	}

	s := NewTradingSerivce(sessionDAO, tradingDAO)

	token := "token1122"
	tradingId := "tradingId1"
	r := s.CreateItem(token, tradingId, "subject", "M/M", "Memo",
		1, 100, 2, 1)
	if r == nil {
		t.Errorf("Result must not be nil")
		return
	}
	if r.Status() != 201 {
		t.Errorf("Wrong status : %d", r.Status())
		return
	}

	json := json(r)
	if v, _ := json.String("id"); v != "item2233" {
		t.Errorf("Wrong id : %s", v)
	}
}

func TestTrading0400_UpdateItem(t *testing.T) {
	sessionDAO := &m.SessionDAO{
		GetByTokenResult: &model.Session{
			Token: "testToken",
		},
	}
	tradingDAO := &m.TradingDAO{
		UpdateItemResult: &model.TradingItem{
			Id: "item2233",
		},
	}

	s := NewTradingSerivce(sessionDAO, tradingDAO)

	token := "token1122"
	id := "item1122"
	tradingId := "tradingId1"
	r := s.UpdateItem(token, id, tradingId, "subject", "M/M", "Memo",
		1, 100, 2, 1)
	if r == nil {
		t.Errorf("Result must not be nil")
		return
	}
	if r.Status() != 200 {
		t.Errorf("Wrong status : %d", r.Status())
		return
	}

	json := json(r)
	if v, _ := json.String("id"); v != "item2233" {
		t.Errorf("Wrong id : %s", v)
	}
}

func TestTrading0500_DeleteItem(t *testing.T) {
	sessionDAO := &m.SessionDAO{
		GetByTokenResult: &model.Session{
			Token: "testToken",
		},
	}
	tradingDAO := &m.TradingDAO{
		SoftDeleteItemResult: nil,
	}

	s := NewTradingSerivce(sessionDAO, tradingDAO)

	token := "token1122"
	id := "item1122"
	tradingId := "tradingId1"
	r := s.DeleteItem(token, id, tradingId)
	if r == nil {
		t.Errorf("Result must not be nil")
		return
	}
	if r.Status() != 204 {
		t.Errorf("Wrong status : %d", r.Status())
		return
	}
}
