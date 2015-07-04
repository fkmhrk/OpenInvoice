package impl

import (
	"../"
	rj "github.com/fkmhrk-go/rawjson"
	"testing"
)

func json(r service.Result) rj.RawJsonObject {
	j, _ := rj.ObjectFromString(r.Body())
	return j
}

func assertString(t *testing.T, item rj.RawJsonObject, key, expected string) {
	if v, _ := item.String(key); v != expected {
		t.Errorf("%s must be %s but %s", key, expected, v)
	}
}

func assertInt(t *testing.T, item rj.RawJsonObject, key string, expected int) {
	if v, _ := item.Int(key); v != expected {
		t.Errorf("%s must be %d but %d", key, expected, v)
	}
}

func assertLong(t *testing.T, item rj.RawJsonObject, key string, expected int64) {
	if v, _ := item.Long(key); v != expected {
		t.Errorf("%s must be %d but %d", key, expected, v)
	}
}

func assertFloat(t *testing.T, item rj.RawJsonObject, key string, expected float64) {
	if v, _ := item.Float(key); v != expected {
		t.Errorf("%s must be %f but %f", key, expected, v)
	}
}

func assertBool(t *testing.T, item rj.RawJsonObject, key string, expected bool) {
	if v, _ := item.Bool(key); v != expected {
		t.Errorf("%s must be %d but %d", key, expected, v)
	}
}
