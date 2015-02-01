package rest

import (
	"bytes"
	"net/http"
)

func readBody(req *http.Request) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)
	return buf.String()
}
