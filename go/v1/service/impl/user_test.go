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
		Id:   "testUser",
		Role: m.Role("Read,Write"),
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
	assertBool(t, json, "isAdmin", false)
}

func TestUser0001_GetToken_Admin(t *testing.T) {
	models := mock.NewMock()
	userDAO, _ := models.User.(*mock.UserDAO)
	userDAO.GetByNamePasswordResult = &m.User{
		Id:   "testUser",
		Role: m.Role("Admin,Read,Write"),
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
	assertBool(t, json, "is_admin", true)
}

func TestUser0100_RefreshToken(t *testing.T) {
	models := mock.NewMock()
	sessionRefreshDAO, _ := models.SessionRefresh.(*mock.SessionRefreshDAO)
	sessionRefreshDAO.GetResult = m.SessionRefresh{
		Token:  "tokenRefresh",
		Role:   "Admin",
		UserId: "user1122",
	}
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	sessionDAO.CreateResult = &m.Session{
		Token: "testToken",
	}

	service := NewUserSerivce(models.User, sessionDAO, models)

	token := "token1122"
	r := service.RefreshToken(token)
	if r == nil {
		t.Errorf("Result must not be nil")
		return
	}
	if r.Status() != 200 {
		t.Errorf("Wrong status : %d", r.Status())
		return
	}
	json := json(r)
	assertString(t, json, "id", "user1122")
	assertString(t, json, "token_type", "bearer")
	assertString(t, json, "access_token", "testToken")
	assertBool(t, json, "is_admin", true)
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

	s := NewUserSerivce(userDAO, sessionDAO, models)

	r := s.GetUsers()
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
	assertString(t, item, "id", "user0")
	assertString(t, item, "display_name", "name0")
	item, _ = userList.Object(1)
	assertString(t, item, "id", "user1")
	assertString(t, item, "display_name", "name1")
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
		Id:          "id1234",
		DisplayName: "loginName",
	}

	s := NewUserSerivce(userDAO, sessionDAO, models)

	token := "token1122"
	r := s.Create(token, "loginName", "disp", "08011112222", "pass1122")
	if r == nil {
		t.Errorf("Result must not be nil")
		return
	}
	if r.Status() != 201 {
		t.Errorf("Wrong status : %d", r.Status())
		return
	}
	json := json(r)
	assertString(t, json, "id", "id1234")
	assertString(t, json, "display_name", "loginName")
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
	r := s.Create(token, "loginName", "disp", "08011112222", "pass1122")
	if r == nil {
		t.Errorf("Result must not be nil")
		return
	}
	if r.Status() != 403 {
		t.Errorf("Wrong status : %d", r.Status())
		return
	}
	json := json(r)
	if v, _ := json.String("msg"); v != MSG_NOT_AUTHORIZED {
		t.Errorf("Wrong msg : %s", v)
	}
}

func TestUser0300_Update(t *testing.T) {
	models := mock.NewMock()
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	sessionDAO.GetByTokenResult = &m.Session{
		Token: "testToken",
		Role:  m.Role("Admin,Read,Write"),
	}
	userDAO, _ := models.User.(*mock.UserDAO)
	userDAO.UpdateResult = &m.User{
		Id: "id1234",
	}

	s := NewUserSerivce(userDAO, sessionDAO, models)

	token := "token1122"
	r := s.Update(token, "user1111", "loginName", "disp", "08011112222", "pass1122")
	if r == nil {
		t.Errorf("Result must not be nil")
		return
	}
	if r.Status() != 200 {
		t.Errorf("Wrong status : %d", r.Status())
		return
	}
	json := json(r)
	assertString(t, json, "id", "user1111")
	assertString(t, json, "login_name", "loginName")
	assertString(t, json, "display_name", "disp")
	assertString(t, json, "tel", "08011112222")
}

func TestUser0301_Update_Not_Admin(t *testing.T) {
	models := mock.NewMock()
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	sessionDAO.GetByTokenResult = &m.Session{
		Token:  "testToken",
		UserId: "user1111",
		Role:   m.Role("Read,Write"),
	}
	userDAO, _ := models.User.(*mock.UserDAO)
	userDAO.UpdateResult = &m.User{
		Id: "id1234",
	}

	s := NewUserSerivce(userDAO, sessionDAO, models)

	token := "token1122"
	r := s.Update(token, "user1111", "loginName", "disp", "08011112222", "pass1122")
	if r == nil {
		t.Errorf("Result must not be nil")
		return
	}
	if r.Status() != 200 {
		t.Errorf("Wrong status : %d", r.Status())
		return
	}
	json := json(r)
	assertString(t, json, "id", "user1111")
	assertString(t, json, "login_name", "loginName")
	assertString(t, json, "display_name", "disp")
	assertString(t, json, "tel", "08011112222")
}

func TestUser0302_Update_Not_Admin_Other(t *testing.T) {
	models := mock.NewMock()
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	sessionDAO.GetByTokenResult = &m.Session{
		Token:  "testToken",
		UserId: "user2222",
		Role:   m.Role("Read,Write"),
	}
	userDAO, _ := models.User.(*mock.UserDAO)
	userDAO.UpdateResult = &m.User{
		Id: "id1234",
	}

	service := NewUserSerivce(userDAO, sessionDAO, models)

	token := "token1122"
	r := service.Update(token, "user1111", "loginName", "disp", "08011112222", "pass1122")
	if r == nil {
		t.Errorf("Result must not be nil")
		return
	}
	if r.Status() != 403 {
		t.Errorf("Wrong status : %d", r.Status())
		return
	}
	json := json(r)
	assertString(t, json, "msg", MSG_NOT_AUTHORIZED)
}

func TestUser0400_Delete(t *testing.T) {
	models := mock.NewMock()
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	sessionDAO.GetByTokenResult = &m.Session{
		Token: "testToken",
		Role:  m.Role("Admin,Read,Write"),
	}
	userDAO, _ := models.User.(*mock.UserDAO)
	userDAO.GetByIdResult = &m.User{
		Id:   "id1234",
		Role: m.Role("Read"),
	}
	userDAO.DeleteResult = nil

	service := NewUserSerivce(userDAO, sessionDAO, models)

	token := "token1122"
	r := service.Delete(token, "user1111")
	if r == nil {
		t.Errorf("Result must not be nil")
		return
	}
	if r.Status() != 204 {
		t.Errorf("Wrong status : %d", r.Status())
		return
	}
	if len(r.Body()) != 0 {
		t.Errorf("Body must be empty but %s", r.Body())
		return
	}
}

func TestUser0401_Delete_Not_Admin(t *testing.T) {
	models := mock.NewMock()
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	sessionDAO.GetByTokenResult = &m.Session{
		Token: "testToken",
		Role:  m.Role("Read,Write"),
	}
	userDAO, _ := models.User.(*mock.UserDAO)
	userDAO.GetByIdResult = &m.User{
		Id:   "id1234",
		Role: m.Role("Read"),
	}
	userDAO.DeleteResult = nil

	service := NewUserSerivce(userDAO, sessionDAO, models)

	token := "token1122"
	r := service.Delete(token, "user1111")
	if r == nil {
		t.Errorf("Result must not be nil")
		return
	}
	if r.Status() != 403 {
		t.Errorf("Wrong status : %d", r.Status())
		return
	}
	json := json(r)
	assertString(t, json, "msg", MSG_NOT_AUTHORIZED)
}

func TestUser0402_Delete_Target_Admin(t *testing.T) {
	models := mock.NewMock()
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	sessionDAO.GetByTokenResult = &m.Session{
		Token: "testToken",
		Role:  m.Role("Admin,Read,Write"),
	}
	userDAO, _ := models.User.(*mock.UserDAO)
	userDAO.GetByIdResult = &m.User{
		Id:   "id1234",
		Role: m.Role("Admin,Read"),
	}
	userDAO.DeleteResult = nil

	service := NewUserSerivce(userDAO, sessionDAO, models)

	token := "token1122"
	r := service.Delete(token, "user1111")
	if r == nil {
		t.Errorf("Result must not be nil")
		return
	}
	if r.Status() != 403 {
		t.Errorf("Wrong status : %d", r.Status())
		return
	}
	json := json(r)
	assertString(t, json, "msg", MSG_NOT_AUTHORIZED)
}
