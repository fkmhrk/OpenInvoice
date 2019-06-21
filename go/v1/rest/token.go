package rest

import (
	"net/http"

	rj "github.com/fkmhrk-go/rawjson"
	s "github.com/fkmhrk/OpenInvoice/v1/service"
	"github.com/mokelab-go/hop"
)

func getToken(user s.UserService) http.HandlerFunc {
	return makeBaseHandler(func(req *http.Request) s.Result {
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
		return user.GetToken(username, password)
	})
}

func refreshToken(services s.Services) http.HandlerFunc {
	return makeBaseHandler(func(req *http.Request) s.Result {
		// read input
		json := rj.RawJsonObject(hop.BodyJSON(req.Context()))
		// get values
		token, _ := json.String("token")
		// execute
		return services.User.RefreshToken(token)
	})
}
