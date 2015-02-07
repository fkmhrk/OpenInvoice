package impl

import (
	"../../model"
	m "../../model/mock"
	"fmt"
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

func TestUser0100_GetList(t *testing.T) {
	var list []*model.User
	for i := 0; i < 2; i++ {
		list = append(list, &model.User{
			Id:          fmt.Sprintf("user%d", i),
			LoginName:   fmt.Sprintf("login%d", i),
			DisplayName: fmt.Sprintf("name%d", i),
		})
	}
	userDAO := &m.UserDAO{
		GetListResult: list,
	}
	sessionDAO := &m.SessionDAO{
		GetByTokenResult: &model.Session{
			Token: "testToken",
		},
	}

	s := NewUserSerivce(userDAO, sessionDAO)

	token := "token1122"
	r := s.GetUsers(token)
	if r == nil {
		t.Errorf("Result must not be nil")
		return
	}
	if r.Status() != 200 {
		t.Errorf("Wrong status : %d", r.Status())
		return
	}
	json := json(r)
	userList, _ := json.Array("users")
	if len(userList) != 2 {
		t.Errorf("Wrong length : %d", len(userList))
	}
	item, _ := userList.Object(0)
	if v, _ := item.String("id"); v != "user0" {
		t.Errorf("Wrong id : %s", v)
	}
	if v, _ := item.String("display_name"); v != "name0" {
		t.Errorf("Wrong display_name : %s", v)
	}
	item, _ = userList.Object(1)
	if v, _ := item.String("id"); v != "user1" {
		t.Errorf("Wrong id : %s", v)
	}
	if v, _ := item.String("display_name"); v != "name1" {
		t.Errorf("Wrong display_name : %s", v)
	}
}
