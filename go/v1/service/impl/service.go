package impl

import (
	s "../"
	m "../../model"
)

func NewServices(models *m.Models) s.Services {
	return s.Services{
		Admin: NewAdminService(models),
	}
}