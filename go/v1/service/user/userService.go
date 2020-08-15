package user

import (
	"net/http"

	m "github.com/fkmhrk/OpenInvoice/v1/model"
	"github.com/fkmhrk/OpenInvoice/v1/model/response"
	"github.com/fkmhrk/OpenInvoice/v1/model/session"
	"github.com/fkmhrk/OpenInvoice/v1/model/user"
	"github.com/fkmhrk/OpenInvoice/v1/rest/service"
	"github.com/mokelab-go/server/entity"
)

type userService struct {
	userDAO           user.DAO
	sessionDAO        session.SessionDAO
	sessionRefreshDAO session.SessionRefreshDAO
}

const (
	TIME_30MIN = 30 * 60
)

func New(u user.DAO, s session.SessionDAO, models *m.Models) service.User {
	return &userService{
		userDAO:           u,
		sessionDAO:        s,
		sessionRefreshDAO: models.SessionRefresh,
	}
}

func (o *userService) GetToken(name, pass string) entity.Response {
	// input check
	if isEmpty(name) {
		return response.Error(http.StatusBadRequest, response.MSG_ERR_NAME_EMPTY)
	}
	if isEmpty(pass) {
		return response.Error(http.StatusBadRequest, response.MSG_ERR_PASS_EMPTY)
	}
	// get User
	user, err := o.userDAO.GetByNamePassword(name, pass)
	if err != nil {
		return response.Error(http.StatusInternalServerError, response.MSG_SERVER_ERROR)
	}
	if user == nil {
		return response.Error(http.StatusBadRequest, response.MSG_WRONG_IDENTIFIER)
	}
	// create session
	session, err := o.sessionDAO.Create(user.Id, string(user.Role), TIME_30MIN)
	if err != nil {
		return response.Error(http.StatusInternalServerError, response.MSG_SERVER_ERROR)
	}
	// create refresh token
	sessionRefresh, err := o.sessionRefreshDAO.Create(user.Id, string(user.Role))
	if err != nil {
		return response.Error(http.StatusInternalServerError, response.MSG_SERVER_ERROR)
	}
	body := map[string]interface{}{
		"id":            user.Id,
		"access_token":  session.Token,
		"refresh_token": sessionRefresh.Token,
		"is_admin":      user.Role.IsAdmin(),
		"token_type":    "bearer",
	}
	return entity.Response{
		Status: http.StatusOK,
		Body:   body,
	}
}

func (o *userService) RefreshToken(token string) entity.Response {
	// input check
	if isEmpty(token) {
		return response.Error(http.StatusBadRequest, response.MSG_TOKEN_EMPTY)
	}
	// get refresh token
	sessionRefresh, err := o.sessionRefreshDAO.Get(token)
	if err != nil {
		return response.Error(http.StatusInternalServerError, response.MSG_SERVER_ERROR)
	}
	// create access token
	session, err := o.sessionDAO.Create(sessionRefresh.UserId, string(sessionRefresh.Role), TIME_30MIN)
	if err != nil {
		return response.Error(http.StatusInternalServerError, response.MSG_SERVER_ERROR)
	}

	body := map[string]interface{}{
		"id":           sessionRefresh.UserId,
		"access_token": session.Token,
		"is_admin":     sessionRefresh.Role.IsAdmin(),
		"token_type":   "bearer",
	}

	return entity.Response{
		Status: http.StatusOK,
		Body:   body,
	}
}

func (s *userService) GetUsers() entity.Response {
	// get User
	users, err := s.userDAO.GetList()
	if err != nil {
		return response.Error(http.StatusInternalServerError, response.MSG_SERVER_ERROR)
	}
	list := make([]interface{}, 0)
	for _, t := range users {
		list = append(list, map[string]interface{}{
			"id":           t.Id,
			"login_name":   t.LoginName,
			"display_name": t.DisplayName,
			"tel":          t.Tel,
		})
	}
	body := map[string]interface{}{
		"users": list,
	}
	return entity.Response{
		Status: http.StatusOK,
		Body:   body,
	}
}

func (s *userService) Create(session *session.Session, loginName, displayName, tel, password string) entity.Response {
	if !session.Role.IsAdmin() {
		return response.Error(http.StatusForbidden, response.MSG_NOT_AUTHORIZED)
	}
	// create
	user, err := s.userDAO.Create(loginName, displayName, "Read,Write", tel, password)
	if err != nil {
		return response.Error(http.StatusInternalServerError, response.MSG_SERVER_ERROR)
	}

	body := map[string]interface{}{
		"id":           user.Id,
		"login_name":   user.LoginName,
		"display_name": user.DisplayName,
		"tel":          user.Tel,
	}
	return entity.Response{
		Status: http.StatusCreated,
		Body:   body,
	}
}

func (s *userService) Update(session *session.Session, id, loginName, displayName, tel, password string) entity.Response {
	if !session.Role.IsAdmin() {
		if session.UserId != id {
			return response.Error(http.StatusForbidden, response.MSG_NOT_AUTHORIZED)
		}
	}
	// update
	_, err := s.userDAO.Update(id, loginName, displayName, "", tel, password)
	if err != nil {
		return response.Error(http.StatusInternalServerError, response.MSG_SERVER_ERROR)
	}

	body := map[string]interface{}{
		"id":           id,
		"login_name":   loginName,
		"display_name": displayName,
		"tel":          tel,
	}
	return entity.Response{
		Status: http.StatusOK,
		Body:   body,
	}
}

func (o *userService) Delete(session *session.Session, id string) entity.Response {
	if !session.Role.IsAdmin() {
		return response.Error(http.StatusForbidden, response.MSG_NOT_AUTHORIZED)
	}
	// get user
	user, err := o.userDAO.GetById(id)
	if err != nil {
		return response.Error(http.StatusInternalServerError, response.MSG_SERVER_ERROR)
	}
	if user.Role.IsAdmin() {
		return response.Error(http.StatusForbidden, response.MSG_NOT_AUTHORIZED)
	}
	// delete
	err = o.userDAO.Delete(id)
	if err != nil {
		return response.Error(http.StatusInternalServerError, response.MSG_SERVER_ERROR)
	}
	return entity.Response{
		Status: 204,
	}
}

func isEmpty(v string) bool {
	return len(v) == 0
}
