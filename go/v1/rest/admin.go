package rest

import (
	m "../model"
	s "../service"
	rj "github.com/fkmhrk-go/rawjson"
	"net/http"
)

func getEnvironment(services s.Services) handler {
	return makeHandler(func(token, tokenType string,
		req *http.Request) s.Result {
		return services.Admin.GetEnvironment(token)
	})
}

func saveEnvironment(services s.Services) handler {
	return makeJsonHandler(func(token, tokenType string, json rj.RawJsonObject) s.Result {
		return services.Admin.SaveEnvironment(token, toEnvList(json))
	})
}

func toEnvList(json rj.RawJsonObject) []*m.Env {
	list := make([]*m.Env, 0, len(json))
	for key, value := range json {
		if strValue, ok := value.(string); ok {
			list = append(list, &m.Env{
				Key:   key,
				Value: strValue,
			})
		}
	}
	return list
}
