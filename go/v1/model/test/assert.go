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

// AssertInt is assert function
func AssertInt(t *testing.T, item map[string]interface{}, key string, expected int) bool {
	actual, ok := item[key].(int)
	if !ok {
		t.Errorf("%s is not int", key)
		return true
	}
	if actual != expected {
		t.Errorf("%s must be %s but %s", key, expected, actual)
		return true
	}
	return false
}

// AssertLong is assert function
func AssertLong(t *testing.T, item map[string]interface{}, key string, expected int64) bool {
	actual, ok := item[key].(int64)
	if !ok {
		t.Errorf("%s is not long", key)
		return true
	}
	if actual != expected {
		t.Errorf("%s must be %s but %s", key, expected, actual)
		return true
	}
	return false
}

// AssertFloat is assert function
func AssertFloat(t *testing.T, item map[string]interface{}, key string, expected float32) bool {
	actual, ok := item[key].(float32)
	if !ok {
		t.Errorf("%s is not float32", key)
		return true
	}
	if actual != expected {
		t.Errorf("%s must be %s but %s", key, expected, actual)
		return true
	}
	return false
}
