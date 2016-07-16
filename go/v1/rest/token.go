package rest

import (
	s "../service"
	rj "github.com/fkmhrk-go/rawjson"
	"net/http"
)

func getToken(user s.UserService) handler {
	return makeBaseHandler(func(req *http.Request) s.Result {
		var username, password string
		contentType := getContentType(req)
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

func refreshToken(services s.Services) handler {
	return makeBaseHandler(func(req *http.Request) s.Result {
		// read input
		json, _ := rj.ObjectFromString(readBody(req))
		// get values
		token, _ := json.String("token")
		// execute
		return services.User.RefreshToken(token)
	})
}
