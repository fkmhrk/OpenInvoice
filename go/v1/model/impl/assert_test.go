package impl

import (
	m "../"
	"fmt"
	"runtime"
	"strings"
	"testing"
)

func assertTradingItem(t *testing.T, item *m.TradingItem,
	id, tradingId, subject string, sortOrder, unitPrice, amount int,
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
		t.Errorf("%s Amount must be %d but %d", caller, amount, item.Amount)
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
