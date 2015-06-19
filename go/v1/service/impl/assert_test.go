package impl

import (
	"fmt"
	rj "github.com/fkmhrk-go/rawjson"
	"runtime"
	"strings"
	"testing"
)

func assertCompany(t *testing.T, item rj.RawJsonObject, id, name, zip, address, phone, unit string) {
	caller := getCaller()
	if v, _ := item.String("id"); v != id {
		t.Errorf("%s id must be %s but %s", caller, id, v)
	}
	if v, _ := item.String("name"); v != name {
		t.Errorf("%s name must be %s but %s", caller, name, v)
	}
	if v, _ := item.String("zip"); v != zip {
		t.Errorf("%s zip must be %s but %s", caller, zip, v)
	}
	if v, _ := item.String("address"); v != address {
		t.Errorf("%s address must be %s but %s", caller, address, v)
	}
	if v, _ := item.String("phone"); v != phone {
		t.Errorf("%s phone must be %s but %s", caller, phone, v)
	}
	if v, _ := item.String("unit"); v != unit {
		t.Errorf("%s unit must be %s but %s", caller, unit, v)
	}
}

func assertTrading(t *testing.T, item rj.RawJsonObject,
	id, companyId, subject string, titleType int,
	workFrom, workTo, total, quotationDate int64, quotationNumber string,
	billDate int64, billNumber string,
	deliveryDate int64, deliveryNumber string,
	taxRate float64, assignee, product, memo string) {
	assertString(t, item, "id", id)
	assertString(t, item, "company_id", companyId)
	assertString(t, item, "subject", subject)
	assertInt(t, item, "title_type", titleType)
	assertLong(t, item, "work_from", workFrom)
	assertLong(t, item, "work_to", workTo)
	assertLong(t, item, "total", total)
	assertLong(t, item, "quotation_date", quotationDate)
	assertString(t, item, "quotation_number", quotationNumber)
	assertLong(t, item, "bill_date", billDate)
	assertString(t, item, "bill_number", billNumber)
	assertLong(t, item, "delivery_date", deliveryDate)
	assertString(t, item, "delivery_number", deliveryNumber)
	assertFloat(t, item, "tax_rate", taxRate)
	assertString(t, item, "assignee", assignee)
	assertString(t, item, "product", product)
	assertString(t, item, "memo", memo)
}

func getCaller() string {
	_, caller, line, _ := runtime.Caller(2)
	path := strings.SplitN(caller, "/", -1)
	return fmt.Sprintf("%s:%d", path[len(path)-1], line)
}
