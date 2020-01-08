package rest

import (
	"fmt"
	"net/http"

	ses "github.com/fkmhrk/OpenInvoice/v1/model/session"
	"github.com/mokelab-go/hop"
)

func getSession(sessionDAO ses.SessionDAO) hop.Op {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			c := r.Context()
			cred := hop.Credential(c)

			session, err := sessionDAO.GetByToken(cred.Token)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "{\"error_code\":\"SERVER_ERROR\",\"msg\":\"Server Error\"}")
				return
			}
			if session == nil {
				w.WriteHeader(401)
				fmt.Fprintf(w, "{\"error_code\":\"WRONG_TOKEN\",\"msg\":\"Wrong token\"}")
				return
			}

			c = setSession(c, session)
			next(w, r.WithContext(c))
		}
	}
}
