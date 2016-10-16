package rest

import (
	m "../model"
	s "../service"
	rj "github.com/fkmhrk-go/rawjson"
	"github.com/mokelab-go/hop"
	"net/http"
)

func getEnvironment(services s.Services) http.HandlerFunc {
	return makeBaseHandler(func(req *http.Request) s.Result {
		return services.Admin.GetEnvironment()
	})
}

func saveEnvironment(services s.Services) http.HandlerFunc {
	return makeBaseHandler(func(req *http.Request) s.Result {
		json := rj.RawJsonObject(hop.BodyJSON(req.Context()))
		return services.Admin.SaveEnvironment(toEnvList(json))
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

func getMyCompanyName(services s.Services) http.HandlerFunc {
	return makeBaseHandler(func(req *http.Request) s.Result {
		return services.Admin.GetMyCompanyname()
	})
}
