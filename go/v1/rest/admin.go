package rest

import (
	s "../service"
	"net/http"
)

func getEnvironment(services s.Services) handler {
	return makeHandler(func(token, tokenType string,
		req *http.Request) s.Result {
		return services.Admin.GetEnvironment(token)
	})
}
