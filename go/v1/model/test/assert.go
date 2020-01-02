package test

import "testing"

// AssertString is assert function
func AssertString(t *testing.T, item map[string]interface{}, key, expected string) bool {
	actual, ok := item[key].(string)
	if !ok {
		t.Errorf("%s is not string", key)
		return true
	}
	if actual != expected {
		t.Errorf("%s must be %s but %s", key, expected, actual)
		return true
	}
	return false
}

// AssertBool is assert function
func AssertBool(t *testing.T, item map[string]interface{}, key string, expected bool) bool {
	actual, ok := item[key].(bool)
	if !ok {
		t.Errorf("%s is not bool", key)
		return true
	}
	if actual != expected {
		t.Errorf("%s must be %s but %s", key, expected, actual)
		return true
	}
	return false
}
