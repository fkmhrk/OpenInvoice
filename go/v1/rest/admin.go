package rest

import (
	"net/http"

	rj "github.com/fkmhrk-go/rawjson"
	"github.com/fkmhrk/OpenInvoice/v1/model/env"
	"github.com/fkmhrk/OpenInvoice/v1/rest/service"
	"github.com/mokelab-go/hop"
)

func getEnvironment(services service.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		resp := services.Admin.GetEnvironment()
		resp.Write(w)
	}
}

func saveEnvironment(services service.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		json := rj.RawJsonObject(hop.BodyJSON(req.Context()))
		resp := services.Admin.SaveEnvironment(toEnvList(json))
		resp.Write(w)
	}
}

func toEnvList(json rj.RawJsonObject) []*env.Env {
	list := make([]*env.Env, 0, len(json))
	for key, value := range json {
		if strValue, ok := value.(string); ok {
			list = append(list, &env.Env{
				Key:   key,
				Value: strValue,
			})
		}
	}
	return list
}

func getMyCompanyName(services service.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		resp := services.Admin.GetMyCompanyname()
		resp.Write(w)
	}
}
