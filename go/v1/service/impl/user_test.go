package impl

import (
	"../../model"
	m "../../model/mock"
	"testing"
)

func TestUser0000_GetToken(t *testing.T) {
	userDAO := &m.UserDAO{
		GetByNamePasswordResult: &model.User{
			Id: "testUser",
		},
	}
	sessionDAO := &m.SessionDAO{
		CreateResult: &model.Session{
			Token: "testToken",
		},
	}

	s := NewUserSerivce(userDAO, sessionDAO)

	name := "user1122"
	pass := "pass2233"
	r := s.GetToken(name, pass)
	if r == nil {
		t.Errorf("Result must not be nil")
		return
	}
	if r.Status() != 200 {
		t.Errorf("Wrong status : %d", r.Status())
		return
	}
	json := json(r)
	if id, _ := json.String("id"); id != "testUser" {
		t.Errorf("Wrong id : %s", id)
	}
	if v, _ := json.String("access_token"); v != "testToken" {
		t.Errorf("Wrong token : %s", v)
	}
	if v, _ := json.String("token_type"); v != "bearer" {
		t.Errorf("Wrong token type : %s", v)
	}
}
