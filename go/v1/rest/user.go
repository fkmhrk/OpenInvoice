package rest

import (
	s "../service"
	"net/http"
)

func getUsers(user s.UserService) handler {
	return makeHandler(func(token, tType string,
		req *http.Request) s.Result {
		return user.GetUsers(token)
	})
}
