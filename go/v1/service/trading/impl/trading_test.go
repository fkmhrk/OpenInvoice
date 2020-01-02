package impl

import (
	"fmt"
	"testing"

	m "github.com/fkmhrk/OpenInvoice/v1/model"
	mock "github.com/fkmhrk/OpenInvoice/v1/model/mock"
	"github.com/fkmhrk/OpenInvoice/v1/model/test"
	s "github.com/fkmhrk/OpenInvoice/v1/service/trading"
)

func TestTrading0000_GetListByUser(t *testing.T) {
	models := mock.NewMock()
	sessionDAO, _ := models.Session.(*mock.SessionDAO)

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
		DeliveryDate:    300,
		DeliveryNumber:  "C300",
		TaxRate:         8.0,
		AssigneeId:      "user2233",
		Product:         "product",
		Memo:            "memo",
	})
	list = append(list, &m.Trading{
		Id: "trade2222",
	})
	tradingDAO, _ := models.Trading.(*mock.TradingDAO)
	tradingDAO.GetListResult = list

	service := New(sessionDAO, tradingDAO, models)

	r := service.GetListByUser()
	if r.Status != 200 {
		t.Errorf("Wrong status : %d", r.Status)
		return
	}

	tradings, _ := r.Body["tradings"].([]interface{})
	if len(tradings) != 2 {
		t.Errorf("Wrong length : %d", len(tradings))
		return
	}
	item, _ := tradings[0].(map[string]interface{})
	assertTrading(t, item, "trade1111", "company2233",
		"subject3344", 1, 1122, 3344, 1980,
		100, "A100",
		200, "B200",
		300, "C300",
		8.0, "user2233", "product", "memo")
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

	service := New(sessionDAO, tradingDAO, models)

	// params
	session := &m.Session{
		Token:  "testToken",
		UserId: "user2233",
	}

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

	r := service.Create(session, companyId, subject, product, memo, titleType, workFrom, workTo, total, quotationDate, billDate, deliveryDate, taxRate)
	if r.Status != 201 {
		t.Errorf("Wrong status : %d", r.Status)
		return
	}

	if v, _ := r.Body["id"].(string); v != "trade1111" {
		t.Errorf("Wrong id : %s", v)
	}
}

func TestTrading0200_Update(t *testing.T) {
	models := mock.NewMock()
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
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

	service := New(sessionDAO, tradingDAO, models)

	// params
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

	r := service.Update(s.Trading{
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
	if r.Status != 200 {
		t.Errorf("Wrong status : %d", r.Status)
		return
	}

	if v, _ := r.Body["id"].(string); v != "trade1111" {
		t.Errorf("Wrong id : %s", v)
	}
}

func TestTrading0200_GetItemsByTradingId(t *testing.T) {
	models := mock.NewMock()
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	var list []*m.TradingItem
	for i := 0; i < 2; i++ {
		list = append(list, &m.TradingItem{
			Id:        fmt.Sprintf("trade%d", i),
			Subject:   fmt.Sprintf("subject%d", i),
			UnitPrice: i*100 + 100,
			Amount:    float64(i*3 + 3),
			Degree:    fmt.Sprintf("D%d", i),
			TaxType:   1,
			Memo:      fmt.Sprintf("memo%d", i),
		})
	}
	tradingDAO, _ := models.Trading.(*mock.TradingDAO)
	tradingDAO.GetItemsByIdResult = list

	s := New(sessionDAO, tradingDAO, models)

	tradingId := "tradingId1"
	r := s.GetItemListByTradingId(tradingId)
	if r.Status != 200 {
		t.Errorf("Wrong status : %d", r.Status)
		return
	}

	items, _ := r.Body["items"].([]interface{})
	if len(items) != 2 {
		t.Errorf("Wrong length : %d", len(items))
		return
	}
	item, _ := items[0].(map[string]interface{})
	if v, _ := item["id"].(string); v != "trade0" {
		t.Errorf("Wrong id : %s", v)
	}
	if v, _ := item["subject"].(string); v != "subject0" {
		t.Errorf("Wrong subject : %s", v)
	}
	if v, _ := item["unit_price"].(int); v != 100 {
		t.Errorf("Wrong unit_price : %d", v)
	}
	if v, _ := item["amount"].(float64); v != 3 {
		t.Errorf("Wrong amount : %f", v)
	}
	if v, _ := item["degree"].(string); v != "D0" {
		t.Errorf("Wrong degree : %s", v)
	}
	if v, _ := item["tax_type"].(int); v != 1 {
		t.Errorf("Wrong tax_type : %d", v)
	}
	if v, _ := item["memo"].(string); v != "memo0" {
		t.Errorf("Wrong memo : %s", v)
	}
}

func TestTrading0300_CreateItem(t *testing.T) {
	models := mock.NewMock()
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	tradingDAO, _ := models.Trading.(*mock.TradingDAO)
	tradingDAO.CreateItemResult = &m.TradingItem{
		Id: "item2233",
	}

	s := New(sessionDAO, tradingDAO, models)

	tradingId := "tradingId1"
	r := s.CreateItem(tradingId, "subject", "M/M", "Memo",
		1, 100, 2, 1)
	if r.Status != 201 {
		t.Errorf("Wrong status : %d", r.Status)
		return
	}

	if v, _ := r.Body["id"].(string); v != "item2233" {
		t.Errorf("Wrong id : %s", v)
	}
}

func TestTrading0400_UpdateItem(t *testing.T) {
	models := mock.NewMock()
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	tradingDAO, _ := models.Trading.(*mock.TradingDAO)
	tradingDAO.UpdateItemResult = &m.TradingItem{
		Id: "item2233",
	}

	s := New(sessionDAO, tradingDAO, models)

	id := "item1122"
	tradingId := "tradingId1"
	r := s.UpdateItem(id, tradingId, "subject", "M/M", "Memo",
		1, 100, 2, 1)
	if r.Status != 200 {
		t.Errorf("Wrong status : %d", r.Status)
		return
	}

	if v, _ := r.Body["id"].(string); v != "item2233" {
		t.Errorf("Wrong id : %s", v)
	}
}

func TestTrading0500_DeleteItem(t *testing.T) {
	models := mock.NewMock()
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	tradingDAO, _ := models.Trading.(*mock.TradingDAO)
	tradingDAO.SoftDeleteItemResult = nil

	s := New(sessionDAO, tradingDAO, models)

	id := "item1122"
	tradingId := "tradingId1"
	r := s.DeleteItem(id, tradingId)
	if r.Status != 204 {
		t.Errorf("Wrong status : %d", r.Status)
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
	s := New(sessionDAO, tradingDAO, models)

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
		date := int64(1434162098716) // 2015-6-13
		r := s.GetNextNumber(item.Arg, date)
		if r.Status != 200 {
			t.Errorf("Wrong status : %d", r.Status)
			return
		}
		test.AssertInt(t, r.Body, "number", 20150004)
		// args check
		if seqDAO.NextSeqType != item.Expected {
			t.Errorf("SeqType must be %s but %d", item.Label, seqDAO.NextSeqType)
		}
	}
}

func TestTrading0700_Delete(t *testing.T) {
	models := mock.NewMock()
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	tradingDAO, _ := models.Trading.(*mock.TradingDAO)
	tradingDAO.DeleteResult = nil

	s := New(sessionDAO, tradingDAO, models)
	r := s.Delete("trade1122")
	if r.Status != 204 {
		t.Errorf("Wrong status : %d", r.Status)
		return
	}
	if r.Body != nil {
		t.Errorf("Body must be empty but %s", r.Body)
	}
}

func assertTrading(t *testing.T, item map[string]interface{},
	id, companyId, subject string, titleType int,
	workFrom, workTo, total, quotationDate int64, quotationNumber string,
	billDate int64, billNumber string,
	deliveryDate int64, deliveryNumber string,
	taxRate float32, assignee, product, memo string) {
	test.AssertString(t, item, "id", id)
	test.AssertString(t, item, "company_id", companyId)
	test.AssertString(t, item, "subject", subject)
	test.AssertInt(t, item, "title_type", titleType)
	test.AssertLong(t, item, "work_from", workFrom)
	test.AssertLong(t, item, "work_to", workTo)
	test.AssertLong(t, item, "total", total)
	test.AssertLong(t, item, "quotation_date", quotationDate)
	test.AssertString(t, item, "quotation_number", quotationNumber)
	test.AssertLong(t, item, "bill_date", billDate)
	test.AssertString(t, item, "bill_number", billNumber)
	test.AssertLong(t, item, "delivery_date", deliveryDate)
	test.AssertString(t, item, "delivery_number", deliveryNumber)
	test.AssertFloat(t, item, "tax_rate", taxRate)
	test.AssertString(t, item, "assignee", assignee)
	test.AssertString(t, item, "product", product)
	test.AssertString(t, item, "memo", memo)
}
