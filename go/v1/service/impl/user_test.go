package impl

import (
	m "../../model"
	mock "../../model/mock"
	"fmt"
	"testing"
)

func TestUser0000_GetToken(t *testing.T) {
	models := mock.NewMock()
	userDAO, _ := models.User.(*mock.UserDAO)
	userDAO.GetByNamePasswordResult = &m.User{
		Id: "testUser",
	}
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	sessionDAO.CreateResult = &m.Session{
		Token: "testToken",
	}
	sessionRefreshDAO, _ := models.SessionRefresh.(*mock.SessionRefreshDAO)
	sessionRefreshDAO.CreateResult = m.SessionRefresh{
		Token: "tokenRefresh",
	}

	s := NewUserSerivce(userDAO, sessionDAO, models)

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
	assertString(t, json, "id", "testUser")
	assertString(t, json, "token_type", "bearer")
	assertString(t, json, "access_token", "testToken")
	assertString(t, json, "refresh_token", "tokenRefresh")
}

func TestUser0100_GetList(t *testing.T) {
	models := mock.NewMock()
	var list []*m.User
	for i := 0; i < 2; i++ {
		list = append(list, &m.User{
			Id:          fmt.Sprintf("user%d", i),
			LoginName:   fmt.Sprintf("login%d", i),
			DisplayName: fmt.Sprintf("name%d", i),
		})
	}
	userDAO, _ := models.User.(*mock.UserDAO)
	userDAO.GetListResult = list
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	sessionDAO.GetByTokenResult = &m.Session{
		Token: "testToken",
	}

	s := NewUserSerivce(userDAO, sessionDAO, models)

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

func TestUser0200_Create(t *testing.T) {
	models := mock.NewMock()
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	sessionDAO.GetByTokenResult = &m.Session{
		Token: "testToken",
		Role:  m.Role("Admin"),
	}
	userDAO, _ := models.User.(*mock.UserDAO)
	userDAO.CreateResult = &m.User{
		Id: "id1234",
	}

	s := NewUserSerivce(userDAO, sessionDAO, models)

	token := "token1122"
	r := s.Create(token, "loginName", "disp", "Read,Write", "pass1122")
	if r == nil {
		t.Errorf("Result must not be nil")
		return
	}
	if r.Status() != 201 {
		t.Errorf("Wrong status : %d", r.Status())
		return
	}
	json := json(r)
	if v, _ := json.String("id"); v != "id1234" {
		t.Errorf("Wrong id : %s", v)
	}
}

func TestUser0201_Create_Not_Admin(t *testing.T) {
	models := mock.NewMock()
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	sessionDAO.GetByTokenResult = &m.Session{
		Token: "testToken",
		Role:  m.Role("Read,Write"),
	}
	userDAO, _ := models.User.(*mock.UserDAO)
	userDAO.CreateResult = &m.User{
		Id: "id1234",
	}

	s := NewUserSerivce(userDAO, sessionDAO, models)

	token := "token1122"
	r := s.Create(token, "loginName", "disp", "Read,Write", "pass1122")
	if r == nil {
		t.Errorf("Result must not be nil")
		return
	}
	if r.Status() != 401 {
		t.Errorf("Wrong status : %d", r.Status())
		return
	}
	json := json(r)
	if v, _ := json.String("msg"); v != MSG_NOT_AUTHORIZED {
		t.Errorf("Wrong msg : %s", v)
	}
}
