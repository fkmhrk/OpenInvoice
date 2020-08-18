package user

import (
	"fmt"
	"testing"

	"github.com/fkmhrk/OpenInvoice/v1/entity"
	mock "github.com/fkmhrk/OpenInvoice/v1/model/mock"
	"github.com/fkmhrk/OpenInvoice/v1/model/response"
	"github.com/fkmhrk/OpenInvoice/v1/model/test"
)

func TestUser0000_GetToken(t *testing.T) {
	models := mock.NewMock()
	userDAO, _ := models.User.(*mock.UserDAO)
	userDAO.GetByNamePasswordResult = &entity.User{
		Id:   "testUser",
		Role: entity.Role("Read,Write"),
	}
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	sessionDAO.CreateResult = &entity.Session{
		Token: "testToken",
	}
	sessionRefreshDAO, _ := models.SessionRefresh.(*mock.SessionRefreshDAO)
	sessionRefreshDAO.CreateResult = entity.SessionRefresh{
		Token: "tokenRefresh",
	}

	s := New(models)

	name := "user1122"
	pass := "pass2233"
	r := s.GetToken(name, pass)
	if r.Status != 200 {
		t.Errorf("Wrong status : %d", r.Status)
		return
	}
	test.AssertString(t, r.Body, "id", "testUser")
	test.AssertString(t, r.Body, "token_type", "bearer")
	test.AssertString(t, r.Body, "access_token", "testToken")
	test.AssertString(t, r.Body, "refresh_token", "tokenRefresh")
	test.AssertBool(t, r.Body, "is_admin", false)
}

func TestUser0001_GetToken_Admin(t *testing.T) {
	models := mock.NewMock()
	userDAO, _ := models.User.(*mock.UserDAO)
	userDAO.GetByNamePasswordResult = &entity.User{
		Id:   "testUser",
		Role: entity.Role("Admin,Read,Write"),
	}
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	sessionDAO.CreateResult = &entity.Session{
		Token: "testToken",
	}
	sessionRefreshDAO, _ := models.SessionRefresh.(*mock.SessionRefreshDAO)
	sessionRefreshDAO.CreateResult = entity.SessionRefresh{
		Token: "tokenRefresh",
	}

	s := New(models)

	name := "user1122"
	pass := "pass2233"
	r := s.GetToken(name, pass)
	if r.Status != 200 {
		t.Errorf("Wrong status : %d", r.Status)
		return
	}
	test.AssertString(t, r.Body, "id", "testUser")
	test.AssertString(t, r.Body, "token_type", "bearer")
	test.AssertString(t, r.Body, "access_token", "testToken")
	test.AssertString(t, r.Body, "refresh_token", "tokenRefresh")
	test.AssertBool(t, r.Body, "is_admin", true)
}

func TestUser0100_RefreshToken(t *testing.T) {
	models := mock.NewMock()
	sessionRefreshDAO, _ := models.SessionRefresh.(*mock.SessionRefreshDAO)
	sessionRefreshDAO.GetResult = entity.SessionRefresh{
		Token:  "tokenRefresh",
		Role:   "Admin",
		UserId: "user1122",
	}
	sessionDAO, _ := models.Session.(*mock.SessionDAO)
	sessionDAO.CreateResult = &entity.Session{
		Token: "testToken",
	}

	service := New(models)

	token := "token1122"
	r := service.RefreshToken(token)
	if r.Status != 200 {
		t.Errorf("Wrong status : %d", r.Status)
		return
	}
	test.AssertString(t, r.Body, "id", "user1122")
	test.AssertString(t, r.Body, "token_type", "bearer")
	test.AssertString(t, r.Body, "access_token", "testToken")
	test.AssertBool(t, r.Body, "is_admin", true)
}

func TestUser0100_GetList(t *testing.T) {
	models := mock.NewMock()
	var list []*entity.User
	for i := 0; i < 2; i++ {
		list = append(list, &entity.User{
			Id:          fmt.Sprintf("user%d", i),
			LoginName:   fmt.Sprintf("login%d", i),
			DisplayName: fmt.Sprintf("name%d", i),
		})
	}
	userDAO, _ := models.User.(*mock.UserDAO)
	userDAO.GetListResult = list

	s := New(models)

	r := s.GetUsers()
	if r.Status != 200 {
		t.Errorf("Wrong status : %d", r.Status)
		return
	}
	userList, _ := r.Body["users"].([]interface{})
	if len(userList) != 2 {
		t.Errorf("Wrong length : %d", len(userList))
	}
	item, _ := userList[0].(map[string]interface{})
	test.AssertString(t, item, "id", "user0")
	test.AssertString(t, item, "display_name", "name0")
	item, _ = userList[1].(map[string]interface{})
	test.AssertString(t, item, "id", "user1")
	test.AssertString(t, item, "display_name", "name1")
}

func TestUser0200_Create(t *testing.T) {
	models := mock.NewMock()
	userDAO, _ := models.User.(*mock.UserDAO)
	userDAO.CreateResult = &entity.User{
		Id:          "id1234",
		DisplayName: "loginName",
	}

	session := &entity.Session{
		Token: "testToken",
		Role:  entity.Role("Admin"),
	}
	s := New(models)

	r := s.Create(session, "loginName", "disp", "08011112222", "pass1122")
	if r.Status != 201 {
		t.Errorf("Wrong status : %d", r.Status)
		return
	}
	test.AssertString(t, r.Body, "id", "id1234")
	test.AssertString(t, r.Body, "display_name", "loginName")
}

func TestUser0201_Create_Not_Admin(t *testing.T) {
	models := mock.NewMock()
	userDAO, _ := models.User.(*mock.UserDAO)
	userDAO.CreateResult = &entity.User{
		Id: "id1234",
	}

	s := New(models)

	session := &entity.Session{
		Token: "testToken",
		Role:  entity.Role("Read,Write"),
	}
	r := s.Create(session, "loginName", "disp", "08011112222", "pass1122")
	if r.Status != 403 {
		t.Errorf("Wrong status : %d", r.Status)
		return
	}
	if v, _ := r.Body["msg"].(string); v != response.MSG_NOT_AUTHORIZED {
		t.Errorf("Wrong msg : %s", v)
	}
}

func TestUser0300_Update(t *testing.T) {
	models := mock.NewMock()
	userDAO, _ := models.User.(*mock.UserDAO)
	userDAO.UpdateResult = &entity.User{
		Id: "id1234",
	}

	s := New(models)

	session := &entity.Session{
		Token: "testToken",
		Role:  entity.Role("Admin,Read,Write"),
	}
	r := s.Update(session, "user1111", "loginName", "disp", "08011112222", "pass1122")
	if r.Status != 200 {
		t.Errorf("Wrong status : %d", r.Status)
		return
	}
	test.AssertString(t, r.Body, "id", "user1111")
	test.AssertString(t, r.Body, "login_name", "loginName")
	test.AssertString(t, r.Body, "display_name", "disp")
	test.AssertString(t, r.Body, "tel", "08011112222")
}

func TestUser0301_Update_Not_Admin(t *testing.T) {
	models := mock.NewMock()
	userDAO, _ := models.User.(*mock.UserDAO)
	userDAO.UpdateResult = &entity.User{
		Id: "id1234",
	}

	s := New(models)

	session := &entity.Session{
		Token:  "testToken",
		UserId: "user1111",
		Role:   entity.Role("Read,Write"),
	}
	r := s.Update(session, "user1111", "loginName", "disp", "08011112222", "pass1122")
	if r.Status != 200 {
		t.Errorf("Wrong status : %d", r.Status)
		return
	}
	test.AssertString(t, r.Body, "id", "user1111")
	test.AssertString(t, r.Body, "login_name", "loginName")
	test.AssertString(t, r.Body, "display_name", "disp")
	test.AssertString(t, r.Body, "tel", "08011112222")
}

func TestUser0302_Update_Not_Admin_Other(t *testing.T) {
	models := mock.NewMock()
	userDAO, _ := models.User.(*mock.UserDAO)
	userDAO.UpdateResult = &entity.User{
		Id: "id1234",
	}

	service := New(models)

	session := &entity.Session{
		Token:  "testToken",
		UserId: "user2222",
		Role:   entity.Role("Read,Write"),
	}
	r := service.Update(session, "user1111", "loginName", "disp", "08011112222", "pass1122")
	if r.Status != 403 {
		t.Errorf("Wrong status : %d", r.Status)
		return
	}
	test.AssertString(t, r.Body, "msg", response.MSG_NOT_AUTHORIZED)
}

func TestUser0400_Delete(t *testing.T) {
	models := mock.NewMock()
	userDAO, _ := models.User.(*mock.UserDAO)
	userDAO.GetByIdResult = &entity.User{
		Id:   "id1234",
		Role: entity.Role("Read"),
	}
	userDAO.DeleteResult = nil

	service := New(models)

	session := &entity.Session{
		Token: "testToken",
		Role:  entity.Role("Admin,Read,Write"),
	}
	r := service.Delete(session, "user1111")
	if r.Status != 204 {
		t.Errorf("Wrong status : %d", r.Status)
		return
	}
	if len(r.Body) != 0 {
		t.Errorf("Body must be empty but %s", r.Body)
		return
	}
}

func TestUser0401_Delete_Not_Admin(t *testing.T) {
	models := mock.NewMock()
	userDAO, _ := models.User.(*mock.UserDAO)
	userDAO.GetByIdResult = &entity.User{
		Id:   "id1234",
		Role: entity.Role("Read"),
	}
	userDAO.DeleteResult = nil

	service := New(models)

	session := &entity.Session{
		Token: "testToken",
		Role:  entity.Role("Read,Write"),
	}
	r := service.Delete(session, "user1111")
	if r.Status != 403 {
		t.Errorf("Wrong status : %d", r.Status)
		return
	}
	test.AssertString(t, r.Body, "msg", response.MSG_NOT_AUTHORIZED)
}

func TestUser0402_Delete_Target_Admin(t *testing.T) {
	models := mock.NewMock()
	userDAO, _ := models.User.(*mock.UserDAO)
	userDAO.GetByIdResult = &entity.User{
		Id:   "id1234",
		Role: entity.Role("Admin,Read"),
	}
	userDAO.DeleteResult = nil

	service := New(models)

	session := &entity.Session{
		Token: "testToken",
		Role:  entity.Role("Admin,Read,Write"),
	}
	r := service.Delete(session, "user1111")
	if r.Status != 403 {
		t.Errorf("Wrong status : %d", r.Status)
		return
	}
	test.AssertString(t, r.Body, "msg", response.MSG_NOT_AUTHORIZED)
}
