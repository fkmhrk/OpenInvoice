package rest

import (
	s "../service"
	rj "github.com/fkmhrk-go/rawjson"
	"net/http"
)

func getUsers(user s.UserService) handler {
	return makeHandler(func(token, tType string,
		req *http.Request) s.Result {
		return user.GetUsers(token)
	})
}

func createUser(services s.Services) handler {
	return makeJsonHandler(func(token, tType string, json rj.RawJsonObject) s.Result {
		// read input
		loginName, _ := json.String("login_name")
		displayName, _ := json.String("display_name")
		password, _ := json.String("password")

		return services.User.Create(token, loginName, displayName, password)
	})
}
