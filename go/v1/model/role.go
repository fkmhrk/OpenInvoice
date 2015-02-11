package model

import (
	"strings"
)

type Role string

const (
	role_admin = "Admin"
	role_read  = "Read"
	role_write = "Write"
)

func (r Role) IsAdmin() bool {
	return r.isContain(role_admin)
}

func (r Role) CanRead() bool {
	return r.isContain(role_read)
}

func (r Role) CanWrite() bool {
	return r.isContain(role_write)
}

func (r Role) isContain(v string) bool {
	roles := r.parse()
	for _, s := range roles {
		if s == v {
			return true
		}
	}
	return false
}

func (r Role) parse() []string {
	return strings.SplitN(string(r), ",", -1)
}
