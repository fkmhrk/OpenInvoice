package rest

import (
	"net/http"

	rj "github.com/fkmhrk-go/rawjson"
	s "github.com/fkmhrk/OpenInvoice/v1/service"
	"github.com/fkmhrk/OpenInvoice/v1/service/user"
	"github.com/mokelab-go/hop"
)

func getUsers(user user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		resp := user.GetUsers()
		resp.Write(w)
	}
}

func createUser(services s.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// read input
		c := req.Context()
		session := session(c)
		json := rj.RawJsonObject(hop.BodyJSON(c))

		loginName, _ := json.String("login_name")
		displayName, _ := json.String("display_name")
		tel, _ := json.String("tel")
		password, _ := json.String("password")

		resp := services.User.Create(session, loginName, displayName, tel, password)
		resp.Write(w)
	}
}

func updateUser(services s.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		c := req.Context()
		id := hop.PathString(c, "id")
		session := session(c)
		json := rj.RawJsonObject(hop.BodyJSON(c))

		// read input
		loginName, _ := json.String("login_name")
		displayName, _ := json.String("display_name")
		tel, _ := json.String("tel")
		password, _ := json.String("password")

		resp := services.User.Update(session, id, loginName, displayName, tel, password)
		resp.Write(w)

	}
}

func deleteUser(services s.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		c := req.Context()
		id := hop.PathString(c, "id")
		session := session(c)

		resp := services.User.Delete(session, id)
		resp.Write(w)

	}
}
