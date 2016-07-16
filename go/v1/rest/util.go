package rest

import (
	s "../service"
	"bytes"
	"fmt"
	rj "github.com/fkmhrk-go/rawjson"
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

func makeJsonHandler(f func(token, tokenType string, json rj.RawJsonObject) s.Result) func(w http.ResponseWriter, req *http.Request) {
	return makeHandler(func(token, tokenType string, req *http.Request) s.Result {
		// to json
		json, _ := rj.ObjectFromString(readBody(req))
		return f(token, tokenType, json)
	})
}

func makeHandler(f func(token, tokenType string, req *http.Request) s.Result) func(w http.ResponseWriter, req *http.Request) {
	return makeBaseHandler(func(req *http.Request) s.Result {
		authorization := req.Header.Get("authorization")
		tokenType, token := parseAuth(authorization)
		return f(token, tokenType, req)
	})
}

func makeBaseHandler(f func(req *http.Request) s.Result) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		printResult(w, f(req))
	}
}

func printResult(w http.ResponseWriter, result s.Result) {
	for k, v := range result.Headers() {
		w.Header().Set(k, v)
	}

	w.WriteHeader(result.Status())
	fmt.Fprintf(w, result.Body())
}
