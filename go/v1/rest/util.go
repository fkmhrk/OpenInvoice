package rest

import (
	"bytes"
	"net/http"
	"strings"
)

func readBody(req *http.Request) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)
	return buf.String()
}

func parseAuth(value string) (string, string) {
	vals := strings.SplitN(value, " ", -1)
	if len(vals) < 2 {
		return "", ""
	}
	return strings.ToUpper(vals[0]), vals[1]
}

func getContentType(req *http.Request) string {
	contentType := req.Header.Get("Content-Type")
	vals := strings.SplitN(contentType, ";", -1)
	return vals[0]
}
