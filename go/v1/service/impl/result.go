package impl

import (
	j "encoding/json"
)

type result struct {
	status  int
	headers map[string]string
	body    string
}

func (r *result) Status() int {
	return r.status
}

func (r *result) Body() string {
	return r.body
}

func (r *result) Headers() map[string]string {
	return r.headers
}

func jsonResult(status int, msg map[string]interface{}) *result {
	if b, err := j.Marshal(msg); err == nil {
		return &result{
			status: status,
			body:   string(b),
		}
	} else {
		return &result{
			status: status,
			body:   "{}",
		}
	}
}

func errorResult(status int, msg string) *result {
	body := map[string]interface{}{
		"msg": msg,
	}
	return jsonResult(status, body)
}
