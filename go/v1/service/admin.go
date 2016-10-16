package service

import (
	m "../model"
)

type AdminService interface {
	// Gets Environment
	GetEnvironment() Result
	// Saves Environment. If key exists, api updates it.
	SaveEnvironment(list []*m.Env) Result
	// Gets My comanpy name
	GetMyCompanyname() Result
}
