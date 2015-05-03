package rest

import (
	rj "github.com/fkmhrk-go/rawjson"
	"testing"
)

func TestAdmin_toEnvList(t *testing.T) {
	json, _ := rj.ObjectFromString(`{"key1":"value1","key2":"value2"}`)
	list := toEnvList(json)
	if len(list) != 2 {
		t.Errorf("List must be 2 but %d", len(list))
	}
	item1 := list[0]
	if item1.Key == "key1" && item1.Value != "value1" {
		t.Errorf("Invalid key-value key=%s / value=%s", item1.Key, item1.Value)
	}
	if item1.Key == "key2" && item1.Value != "value2" {
		t.Errorf("Invalid key-value key=%s / value=%s", item1.Key, item1.Value)
	}
}
