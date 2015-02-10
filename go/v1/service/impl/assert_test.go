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

func getCaller() string {
	_, caller, line, _ := runtime.Caller(2)
	path := strings.SplitN(caller, "/", -1)
	return fmt.Sprintf("%s:%d", path[len(path)-1], line)
}
