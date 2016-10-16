package rest

import (
	s "../service"
	rj "github.com/fkmhrk-go/rawjson"
	"github.com/mokelab-go/hop"
	"net/http"
)

func getUsers(user s.UserService) http.HandlerFunc {
	return makeBaseHandler(func(req *http.Request) s.Result {
		return user.GetUsers()
	})
}

func createUser(services s.Services) http.HandlerFunc {
	return makeBaseHandler(func(req *http.Request) s.Result {
		// read input
		c := req.Context()
		session := session(c)
		json := rj.RawJsonObject(hop.BodyJSON(c))

		loginName, _ := json.String("login_name")
		displayName, _ := json.String("display_name")
		tel, _ := json.String("tel")
		password, _ := json.String("password")

		return services.User.Create(session, loginName, displayName, tel, password)
	})
}

func updateUser(services s.Services) http.HandlerFunc {
	return makeBaseHandler(func(req *http.Request) s.Result {
		c := req.Context()
		id := hop.PathString(c, "id")
		session := session(c)
		json := rj.RawJsonObject(hop.BodyJSON(c))

		// read input
		loginName, _ := json.String("login_name")
		displayName, _ := json.String("display_name")
		tel, _ := json.String("tel")
		password, _ := json.String("password")

		return services.User.Update(session, id, loginName, displayName, tel, password)
	})
}

func deleteUser(services s.Services) http.HandlerFunc {
	return makeBaseHandler(func(req *http.Request) s.Result {
		c := req.Context()
		id := hop.PathString(c, "id")
		session := session(c)

		return services.User.Delete(session, id)
	})
}
