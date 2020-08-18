package service

import (
	e "github.com/fkmhrk/OpenInvoice/v1/entity"
	"github.com/mokelab-go/server/entity"
)

// Admin is service
type Admin interface {
	// Gets Environment
	GetEnvironment() entity.Response
	// Saves Environment. If key exists, api updates it.
	SaveEnvironment(list []*e.Env) entity.Response
	// Gets My comanpy name
	GetMyCompanyname() entity.Response
}
