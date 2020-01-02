package admin

import (
	m "github.com/fkmhrk/OpenInvoice/v1/model"
	"github.com/mokelab-go/server/entity"
)

type Service interface {
	// Gets Environment
	GetEnvironment() entity.Response
	// Saves Environment. If key exists, api updates it.
	SaveEnvironment(list []*m.Env) entity.Response
	// Gets My comanpy name
	GetMyCompanyname() entity.Response
}
