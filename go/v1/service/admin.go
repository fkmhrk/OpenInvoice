package service

import (
	m "../model"
)

type AdminService interface {
	// Gets Environment
	GetEnvironment(token string) Result
	// Saves Environment. If key exists, api updates it.
	SaveEnvironment(token string, list []*m.Env) Result
}
