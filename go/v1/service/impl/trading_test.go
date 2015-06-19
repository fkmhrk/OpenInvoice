package impl

import (
	s "../"
	m "../../model"
	mock "../../model/mock"
	"fmt"
	"testing"
)

func TestTrading0000_GetListByUser(t *testing.T) {
	models := mock.NewMock()
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	sessionDAO.GetByTokenResult = &m.Session{
		Token: "testToken",
	}

	var list []*m.Trading
	list = append(list, &m.Trading{
		Id:              "trade1111",
		CompanyId:       "company2233",
		TitleType:       1,
		Subject:         "subject3344",
		Total:           1980,
		WorkFrom:        1122,
		WorkTo:          3344,
		QuotationDate:   100,
		QuotationNumber: "A100",
		BillDate:        200,
		BillNumber:      "B200",
		TaxRate:         8.0,
		AssigneeId:      "user2233",
		Product:         "product",
	})
	list = append(list, &m.Trading{
		Id: "trade2222",
	})
	tradingDAO, _ := models.Trading.(*mock.TradingDAO)
	tradingDAO.GetListByUserResult = list

	service := NewTradingSerivce(sessionDAO, tradingDAO, models)

	token := "token1122"
	r := service.GetListByUser(token)
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
	assertTrading(t, item, "trade1111", "company2233",
		"subject3344", 1, 1122, 3344, 1980,
		100, "A100",
		200, "B200",
		8.0, "user2233", "product")
}

func TestTrading0100_Create(t *testing.T) {
	models := mock.NewMock()
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	sessionDAO.GetByTokenResult = &m.Session{
		Token: "testToken",
	}
	tradingDAO, _ := models.Trading.(*mock.TradingDAO)
	tradingDAO.CreateResult = &m.Trading{
		Id:         "trade1111",
		CompanyId:  "company2233",
		Subject:    "subject3344",
		WorkFrom:   1122,
		WorkTo:     3344,
		AssigneeId: "user2233",
		Product:    "product",
	}

	service := NewTradingSerivce(sessionDAO, tradingDAO, models)

	// params
	token := "token1122"
	companyId := "company1122"
	titleType := 1
	subject := "subject3344"
	product := "product4455"
	memo := "memo"
	workFrom := int64(100)
	workTo := int64(200)
	total := int64(1980)
	quotationDate := int64(300)
	billDate := int64(400)
	deliveryDate := int64(500)
	taxRate := float32(8)

	r := service.Create(token, companyId, subject, product, memo, titleType, workFrom, workTo, total, quotationDate, billDate, deliveryDate, taxRate)
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
	models := mock.NewMock()
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	sessionDAO.GetByTokenResult = &m.Session{
		Token:  "testToken",
		UserId: "user1122",
	}
	tradingDAO, _ := models.Trading.(*mock.TradingDAO)
	tradingDAO.GetByIdResult = &m.Trading{
		Id:         "trade1111",
		CompanyId:  "company2233",
		Subject:    "subject3344",
		WorkFrom:   2233,
		WorkTo:     4455,
		AssigneeId: "user2233",
		Product:    "product",
	}

	tradingDAO.UpdateResult = &m.Trading{
		Id:         "trade1111",
		CompanyId:  "company2233",
		Subject:    "subject3344",
		WorkFrom:   2233,
		WorkTo:     4455,
		AssigneeId: "user2233",
		Product:    "product",
	}

	service := NewTradingSerivce(sessionDAO, tradingDAO, models)

	// params
	token := "token1122"
	id := "20150203"
	companyId := "company1122"
	subject := "subject3344"
	product := "product4455"
	titleType := 1
	workFrom := int64(100)
	workTo := int64(200)
	total := int64(1280)
	quotationDate := int64(300)
	billDate := int64(400)
	taxRate := float32(8)

	r := service.Update(token, s.Trading{
		m.Trading{
			Id:            id,
			CompanyId:     companyId,
			Subject:       subject,
			Product:       product,
			TitleType:     titleType,
			WorkFrom:      workFrom,
			WorkTo:        workTo,
			Total:         total,
			QuotationDate: quotationDate,
			BillDate:      billDate,
			TaxRate:       taxRate,
		},
	})
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
	models := mock.NewMock()
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	sessionDAO.GetByTokenResult = &m.Session{
		Token: "testToken",
	}
	var list []*m.TradingItem
	for i := 0; i < 2; i++ {
		list = append(list, &m.TradingItem{
			Id:        fmt.Sprintf("trade%d", i),
			Subject:   fmt.Sprintf("subject%d", i),
			UnitPrice: i*100 + 100,
			Amount:    i*3 + 3,
			Degree:    fmt.Sprintf("D%d", i),
			TaxType:   1,
			Memo:      fmt.Sprintf("memo%d", i),
		})
	}
	tradingDAO, _ := models.Trading.(*mock.TradingDAO)
	tradingDAO.GetItemsByIdResult = list

	s := NewTradingSerivce(sessionDAO, tradingDAO, models)

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
	models := mock.NewMock()
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	sessionDAO.GetByTokenResult = &m.Session{
		Token: "testToken",
	}
	tradingDAO, _ := models.Trading.(*mock.TradingDAO)
	tradingDAO.CreateItemResult = &m.TradingItem{
		Id: "item2233",
	}

	s := NewTradingSerivce(sessionDAO, tradingDAO, models)

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
	models := mock.NewMock()
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	sessionDAO.GetByTokenResult = &m.Session{
		Token: "testToken",
	}
	tradingDAO, _ := models.Trading.(*mock.TradingDAO)
	tradingDAO.UpdateItemResult = &m.TradingItem{
		Id: "item2233",
	}

	s := NewTradingSerivce(sessionDAO, tradingDAO, models)

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
	models := mock.NewMock()
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	sessionDAO.GetByTokenResult = &m.Session{
		Token: "testToken",
	}
	tradingDAO, _ := models.Trading.(*mock.TradingDAO)
	tradingDAO.SoftDeleteItemResult = nil

	s := NewTradingSerivce(sessionDAO, tradingDAO, models)

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

func TestTrading0600_GetNextNumber(t *testing.T) {
	models := mock.NewMock()
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	sessionDAO.GetByTokenResult = &m.Session{
		Token: "testToken",
	}
	envDAO, _ := models.Env.(*mock.EnvDAO)
	envDAO.GetResult = m.Env{
		Value: "3",
	}
	tradingDAO, _ := models.Trading.(*mock.TradingDAO)
	seqDAO, _ := models.Seq.(*mock.SeqDAO)
	seqDAO.NextResult = m.Seq{
		Value: 4,
	}
	s := NewTradingSerivce(sessionDAO, tradingDAO, models)

	table := []struct {
		Arg      string
		Expected m.SeqType
		Label    string
	}{
		{"quotation", m.SeqType_Quotation, "m.SeqType_Quotation"},
		{"delivery", m.SeqType_Delivery, "m.SeqType_Delivery"},
		{"bill", m.SeqType_Bill, "m.SeqType_Bill"},
	}

	for _, item := range table {
		token := "token1122"
		date := int64(1434162098716) // 2015-6-13
		r := s.GetNextNumber(token, item.Arg, date)
		if r == nil {
			t.Errorf("Result must not be nil")
			return
		}
		if r.Status() != 200 {
			t.Errorf("Wrong status : %d", r.Status())
			return
		}
		body := json(r)
		assertInt(t, body, "number", 20150004)
		// args check
		if seqDAO.NextSeqType != item.Expected {
			t.Errorf("SeqType must be %s but %d", item.Label, seqDAO.NextSeqType)
		}
	}
}

func TestTrading0700_Delete(t *testing.T) {
	models := mock.NewMock()
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	sessionDAO.GetByTokenResult = &m.Session{
		Token: "testToken",
	}
	tradingDAO, _ := models.Trading.(*mock.TradingDAO)
	tradingDAO.DeleteResult = nil

	s := NewTradingSerivce(sessionDAO, tradingDAO, models)
	token := "token1122"
	r := s.Delete(token, "trade1122")
	if r == nil {
		t.Errorf("Result must not be nil")
		return
	}
	if r.Status() != 204 {
		t.Errorf("Wrong status : %d", r.Status())
		return
	}
	if r.Body() != "" {
		t.Errorf("Body must be empty but %s", r.Body())
	}
}
