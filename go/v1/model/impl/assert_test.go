package impl

import (
	m "../"
	"fmt"
	"math"
	"runtime"
	"strings"
	"testing"
)

func assertUser(t *testing.T, item *m.User, id, loginName, displayName, role, tel string) {
	caller := getCaller()
	if item.Id != id {
		t.Errorf("%s Id must be %s but %s", caller, id, item.Id)
	}
	if item.LoginName != loginName {
		t.Errorf("%s LoginName must be %s but %s", caller, loginName, item.LoginName)
	}
	if item.DisplayName != displayName {
		t.Errorf("%s DisplayName must be %s but %s", caller,
			displayName, item.DisplayName)
	}
	if string(item.Role) != role {
		t.Errorf("%s Role must be %s but %s", caller,
			role, string(item.Role))
	}
	if item.Tel != tel {
		t.Errorf("[%s] Tel must be %s but %s", caller, tel, item.Tel)
	}
}

func assertCompany(t *testing.T, item *m.Company,
	id, name, zip, address, phone, unit string) {
	caller := getCaller()
	if item.Id != id {
		t.Errorf("%s Id must be %s but %s", caller, id, item.Id)
	}
	if item.Name != name {
		t.Errorf("%s Name must be %s but %s", caller, name, item.Name)
	}
	if item.Zip != zip {
		t.Errorf("%s Zip must be %s but %s", caller, zip, item.Zip)
	}
	if item.Address != address {
		t.Errorf("%s Address must be %s but %s", caller, address, item.Address)
	}
	if item.Phone != phone {
		t.Errorf("%s Phone must be %s but %s", caller, phone, item.Phone)
	}
	if item.Unit != unit {
		t.Errorf("%s Unit must be %s but %s", caller, unit, item.Unit)
	}
}

func assertTrading(t *testing.T, item *m.Trading,
	id, companyId, subject string, titleType int,
	workFrom, workTo, quotationDate, billDate int64,
	taxRate float32, assignee, product string) {
	caller := getCaller()
	if item.Id != id {
		t.Errorf("%s Id must be %s but %s", caller, id, item.Id)
	}
	if item.CompanyId != companyId {
		t.Errorf("%s CompanyId must be %s but %s", caller, companyId, item.CompanyId)
	}
	if item.Subject != subject {
		t.Errorf("%s Subject must be %s but %s", caller, subject, item.Subject)
	}
	if item.TitleType != titleType {
		t.Errorf("%s TitleType must be %d but %d", caller, titleType, item.TitleType)
	}
	if item.WorkFrom != workFrom {
		t.Errorf("%s WorkFrom must be %d but %d", caller, workFrom, item.WorkFrom)
	}
	if item.WorkTo != workTo {
		t.Errorf("%s WorkTo must be %d but %d", caller, workTo, item.WorkTo)
	}
	if item.QuotationDate != quotationDate {
		t.Errorf("%s QuotationDate must be %d but %d", caller, quotationDate, item.QuotationDate)
	}
	if item.BillDate != billDate {
		t.Errorf("%s BillDate must be %d but %d", caller, billDate, item.BillDate)
	}
	if math.Abs(float64(item.TaxRate-taxRate)) > 0.1 {
		t.Errorf("%s TaxRate must be %f but %f", caller, taxRate, item.TaxRate)
	}
	if item.AssigneeId != assignee {
		t.Errorf("%s AssigneeId must be %s but %s", caller, assignee, item.AssigneeId)
	}
	if item.Product != product {
		t.Errorf("%s Product must be %s but %s", caller, product, item.Product)
	}

}

func assertTradingItem(t *testing.T, item *m.TradingItem,
	id, tradingId, subject string, sortOrder, unitPrice int, amount float64,
	degree string, taxType int, memo string) {
	caller := getCaller()
	if item.Id != id {
		t.Errorf("%s Id must be %s but %s", caller, id, item.Id)
	}
	if item.TradingId != tradingId {
		t.Errorf("%s TradingId must be %s but %s", caller, tradingId, item.TradingId)
	}
	if item.Subject != subject {
		t.Errorf("%s Subject must be %s but %s", caller, subject, item.Subject)
	}
	if item.SortOrder != sortOrder {
		t.Errorf("%s SortOrder must be %d but %d", caller, sortOrder, item.SortOrder)
	}
	if item.UnitPrice != unitPrice {
		t.Errorf("%s UnitPrice must be %d but %d", caller, unitPrice, item.UnitPrice)
	}
	if item.Amount != amount {
		t.Errorf("%s Amount must be %f but %f", caller, amount, item.Amount)
	}
	if item.Degree != degree {
		t.Errorf("%s Degree must be %s but %s", caller, degree, item.Degree)
	}
	if item.TaxType != taxType {
		t.Errorf("%s TaxType must be %d but %d", caller, taxType, item.TaxType)
	}
	if item.Memo != memo {
		t.Errorf("%s Memo must be %s but %s", caller, memo, item.Memo)
	}
}

func getCaller() string {
	_, caller, line, _ := runtime.Caller(2)
	path := strings.SplitN(caller, "/", -1)
	return fmt.Sprintf("%s:%d", path[len(path)-1], line)
}
