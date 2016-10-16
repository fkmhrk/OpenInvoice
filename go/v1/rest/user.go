package rest

import (
	s "../service"
	rj "github.com/fkmhrk-go/rawjson"
	"github.com/gorilla/mux"
	"net/http"
)

func getUsers(user s.UserService) http.HandlerFunc {
	return makeBaseHandler(func(req *http.Request) s.Result {
		return user.GetUsers()
	})
}

func createUser(services s.Services) handler {
	return makeJsonHandler(func(token, tType string, json rj.RawJsonObject) s.Result {
		// read input
		loginName, _ := json.String("login_name")
		displayName, _ := json.String("display_name")
		tel, _ := json.String("tel")
		password, _ := json.String("password")

		return services.User.Create(token, loginName, displayName, tel, password)
	})
}

func updateUser(services s.Services) handler {
	return makeHandler(func(token, tType string, req *http.Request) s.Result {
		// read path param
		vars := mux.Vars(req)
		id := vars["id"]
		// to json
		json, _ := rj.ObjectFromString(readBody(req))
		// read input
		loginName, _ := json.String("login_name")
		displayName, _ := json.String("display_name")
		tel, _ := json.String("tel")
		password, _ := json.String("password")

		return services.User.Update(token, id, loginName, displayName, tel, password)
	})
}

func deleteUser(services s.Services) handler {
	return makeHandler(func(token, tType string, req *http.Request) s.Result {
		// read path param
		vars := mux.Vars(req)
		id := vars["id"]

		return services.User.Delete(token, id)
	})
}
