package rest

import (
	"net/http"

	rj "github.com/fkmhrk-go/rawjson"
	"github.com/fkmhrk/OpenInvoice/v1/rest/service"
	"github.com/mokelab-go/hop"
)

func getToken(user service.User) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var username, password string
		contentType := hop.ContentType(req.Context())
		if contentType == "multipart/form-data" {
			req.ParseMultipartForm(8192)
			username = req.Form.Get("username")
			password = req.Form.Get("password")
		} else {
			// read input
			json, _ := rj.ObjectFromString(readBody(req))
			// get values
			username, _ = json.String("username")
			password, _ = json.String("password")
		}
		// execute
		resp := user.GetToken(username, password)
		resp.Write(w)
	}
}

func refreshToken(services service.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// read input
		json := rj.RawJsonObject(hop.BodyJSON(req.Context()))
		// get values
		token, _ := json.String("token")
		// execute
		resp := services.User.RefreshToken(token)
		resp.Write(w)
	}
}
