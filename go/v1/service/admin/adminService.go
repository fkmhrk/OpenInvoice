package admin

import (
	"github.com/fkmhrk/OpenInvoice/v1/model/env"
	"github.com/mokelab-go/server/entity"
)

// Service is
type Service interface {
	// Gets Environment
	GetEnvironment() entity.Response
	// Saves Environment. If key exists, api updates it.
	SaveEnvironment(list []*env.Env) entity.Response
	// Gets My comanpy name
	GetMyCompanyname() entity.Response
}
